package main

import (
  "fmt"
  "flag"
  "strings"
)

// const
const boxPadding = 2

// tree
type Node struct {
  token bool
  text []string
  left, right *Node
}

type Tree struct {
  root *Node
}

func newTree () *Tree {
  return new(Tree)
}

func newNode (text string) *Node {
  n := new(Node)
  n.token = containsToken(text)
  n.text = split(text)
  return n
}

func split (s string) []string {
  return strings.Split(s, "|")
}

func (t *Tree) insertTree (text string) {
  t.root = insertNode(t.root, text)
}

func insertNode (node *Node, text string) *Node {
  switch {
  case node == nil:
    return newNode(text)
  case node.left == nil:
    node.left = insertNode(node.left, text)
  default:
    node.right = insertNode(node.right, text)
  }
  return node
}

func (t *Tree) maxLine () int {
  root, line := t.root, 0
  if root != nil {
    line = len(root.text)

    if root.left != nil {
      maxLineNode(root.left, &line)
    }
    if root.right != nil {
      maxLineNode(root.right, &line)
    }
  }
  return line
}

func maxLineNode (n *Node, line *int) {
  if len(n.text) > *line {
    *line = len(n.text)
  }
  if n.left != nil {
    maxLineNode(n.left, line)
  }
  if n.right != nil {
    maxLineNode(n.right, line)
  }
}



func (t *Tree) printTree () {
  root := t.root
  if root != nil {
    fmt.Println("0", root.token, root.text)
    if root.left != nil {
      printNode(root.left, 1)
    }
    if root.right != nil {
      printNode(root.right, 1)
    }
  }
}

func printNode (n *Node, count int) {
  fmt.Println(count, n.token, n.text)

  count += 1
  if n.left != nil {
    printNode(n.left, count)
  }
  if n.right != nil {
    printNode(n.right, count)
  }
}

// main
func main () {
  flag.Parse()
  text := flag.Args()[0]

  split := parse(text)
  // split.printTree()

  box := createBox(split)
  for _, v := range box {
    fmt.Println(v)
  }
}

// split
func parse (text string) *Tree {
  // var result []string
  tree := newTree()
  str, tbuf, tfound := "", "", false // tbuf => token buffer

  for _, c := range text {
    c := string([]rune{c})

    switch c {
    case "+":
      if tbuf == "" || tbuf == "<-" {
        tfound = true
      }

    case "<":
      if tbuf == "" {
        tfound = true
      }

    case "-":
      if tbuf == "<" || tbuf == "+" {
        tfound = true
      }

    case ">":
      if tbuf == "+-" || tbuf == "<-" {
        tfound = true
      }
    }

    if tfound {
      tfound = false
      tbuf += c
    } else {
      str += tbuf
      str += c
      tbuf = ""
    }

    if len(tbuf) > 2 {
      if containsToken(tbuf) {
        tree.insertTree(tbuf)
        tree.insertTree(strings.TrimSpace(str))
        str, tbuf = "", ""
      } else {
        str += tbuf
      }
    }
  }

  tree.insertTree(strings.TrimSpace(str))
  return tree
}


// ---
func createBox (t *Tree) []string {
  var box []string

  if t != nil {
    constructBox(t.root, &box, t.maxLine()) // line number
  }
  return box
}

func constructBox (n *Node, box *[]string, line int) {
  if n.left != nil {
    constructBox(n.left, box, line)
  }

  if len(*box) < 3 {
    if n.token { // IS TOKEN?
      *box = append(*box, "     ")
      *box = append(*box, " " + n.text[0] + " ")
      *box = append(*box, "     ")
    } else {     // No, I'm not Token
      if len(n.text) > 1 {

        lineLength := 0
        for _, v := range n.text {
          if len(v) > lineLength {
            lineLength = len(v)
          }
        }
        frame := createFrame(lineLength)
        writeLine := calcWriteLine(line, len(n.text))

        for i := 0; i < line + 2; i++ {
          if i == 0 || i == line + 1 {
            *box = append(*box, frame)
          }

          if i - writeLine > -1 && i - writeLine < len(n.text) {
            fmt.Println("f", (lineLength - len(n.text[i - writeLine])) / 2)
            leftPad := boxPadding + ((lineLength - len(n.text[i - writeLine])) / 2)
            rightPad := leftPad
            if lineLength > len(n.text[i - writeLine]) && (lineLength - len(n.text[i - writeLine]) / 2) > 0 {
              rightPad += 1
            }

            *box = append(*box, "|" + strings.Repeat(" ", leftPad) + n.text[i - writeLine] + strings.Repeat(" ", rightPad) +"|")
          }
        }


      } else {
        frame := createFrame(len(n.text[0]))
        pad := strings.Repeat(" ", boxPadding)

        *box = append(*box, frame)
        *box = append(*box, "|" + pad + n.text[0] + pad +"|")
        *box = append(*box, frame)
      }
    }

  } else {
    // 3 LINE OVER

    if n.token { // Token
      (*box)[0] += "     "
      (*box)[1] += " " + n.text[0] + " "
      (*box)[2] += "     "

      if len(n.text) > 1 {
      } else {
      }
    } else {
      frame := createFrame(len(n.text[0]))
      pad := strings.Repeat(" ", boxPadding)

      (*box)[0] += frame
      (*box)[1] += "|" + pad + n.text[0] + pad +"|"
      (*box)[2] += frame
    }
  }

  if n.right != nil {
    constructBox(n.right, box, line)
  }
}

func calcWriteLine (maxLine int, sentenceLine int) int {
  res := 1
  // if sentenceLine * 2 == maxLine {
  //  res += 1
  // }
  return res
}

func createFrame (num int) string {
  frame := ""
  for i := 0; i < num + 2 + boxPadding * 2; i++ {
    if i == 0 || i == num + 1 + boxPadding * 2 {
      frame += "+"
    } else {
      frame += "-"
    }
  }
  return frame
}

func containsToken (t string) bool {
  tokens := [...]string{
    "<-+",
    "+->",
    "<->",
  }

  for _, v := range tokens {
    if t == v {
      return true
    }
  }
  return false
}
