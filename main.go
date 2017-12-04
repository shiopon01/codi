package main

import (
  "fmt"
  "flag"
)


func main () {
  flag.Parse()
  text := flag.Args()[0]

  split := parse(text)
  box := createBox(split)

  for _, v := range box {
    fmt.Println(v)
  }
}

// split
func parse (text string) []string {
  var result []string
  str, tbuf := "", ""

  for _, c := range text {
    c := string([]rune{c})

    switch c {
    case "+":
      if tbuf == "" || tbuf == "<-" {
        tbuf += c
      } else {
        str += tbuf
        str += c
        tbuf = ""
      }

    case "<":
      if tbuf == "" {
        tbuf += c
      } else {
        str += tbuf
        str += c
        tbuf = ""
      }

    case "-":
      if tbuf == "<" || tbuf == "+" {
        tbuf += c
      } else {
        str += tbuf
        str += c
        tbuf = ""
      }

    case ">":
      if tbuf == "+-" || tbuf == "<-" {
        tbuf += c
      } else {
        str += tbuf
        str += c
        tbuf = ""
      }

    default:
      str += tbuf
      str += c
      tbuf = ""
    }

    if len(tbuf) > 2 {
      if containsToken(tbuf) {
        result = append(result, str)
        result = append(result, tbuf)
        str, tbuf = "", ""
      } else {
        str += tbuf
      }
    }
  }

  result = append(result, str)
  return result
}


// ---
func createBox (tokens []string) []string {
  result, line := []string{"", "", ""}, ""

  for _, v := range tokens {
    if containsToken(v) {
      result[0] += "     "
      result[1] += " " + v + " "
      result[2] += "     "

    } else {
      for i := 0; i < len(v) + 4; i++ {
        if i == 0 || i == len(v) + 3 {
          line += "+"
        } else {
          line += "-"
        }
      }

      result[0] += line
      result[1] += "| " + v + " |"
      result[2] += line
      line = ""
    }
  }

  return result
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
