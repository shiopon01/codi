package main

import (
  "fmt"
  "flag"
)

func main () {
  flag.Parse()
  text := flag.Args()[0]

  tokens := parse(text)
  box := createBox(tokens)

  for _, v := range box {
    fmt.Println(v)
  }
}

func parse (text string) []string {
  result := make([]string, 0, 5)
  str, buf := "", ""

  for _, c := range text {
    c := string([]rune{c})

    switch c {
    case "-":
      if buf == "<" || buf == "+" {
        buf += c
      } else {
        str += buf
        buf = ""
      }

    case "+":
      if buf == "" {
        buf += c
      } else if buf == "<-" {
        buf += c
        buf += c
        result = append(result, str)
        result = append(result, buf)
        str, buf = "", ""
      } else {
        str += buf
        buf = ""
      }

    case "<":
      if buf == "" {
        buf += c
      }

    case ">":
      if buf == "+-" || buf == "<-" {
        buf += c
        result = append(result, str)
        result = append(result, buf)
        str, buf = "", ""
      } else {
        str += buf
        buf = ""
      }

    default:
      str += c
    }
  }
  result = append(result, str)

  return result
}


func createBox (tokens []string) []string {
  result, line := []string{"", "", ""}, ""

  for _, v := range tokens {
    if v == "<-+" || v == "<->" || v == "+->" {
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
