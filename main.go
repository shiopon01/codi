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
  text string
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
  n.text = text
  return n
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
        tree.insertTree(strings.TrimSpace(tbuf))
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
    constructBox(t.root, &box)
  }
  return box
}

func constructBox (n *Node, box *[]string) {
  if n.left != nil {
    constructBox(n.left, box)
  }

  if n.token {
    // token
    if len(*box) < 3 {
      *box = append(*box, "     ")
      *box = append(*box, " " + n.text + " ")
      *box = append(*box, "     ")
    } else {
      (*box)[0] += "     "
      (*box)[1] += " " + n.text + " "
      (*box)[2] += "     "
    }
  } else {
    // not token
    pad := strings.Repeat(" ", boxPadding)
    line := ""
    for i := 0; i < len(n.text) + 2 + boxPadding * 2; i++ {
      if i == 0 || i == len(n.text) + 1 + boxPadding * 2 {
        line += "+"
      } else {
        line += "-"
      }
    }

    if len(*box) < 3 {
      *box = append(*box, line)
      *box = append(*box, "|" + pad + n.text + pad +"|")
      *box = append(*box, line)
    } else {
      (*box)[0] += line
      (*box)[1] += "|" + pad + n.text + pad +"|"
      (*box)[2] += line
    }
  }

  if n.right != nil {
    constructBox(n.right, box)
  }
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
