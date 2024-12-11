package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)

var lines []string
func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
    return
  }
  bytes, err := io.ReadAll(file)
  if err != nil {
    fmt.Println(err)
    return
  }
  lines = strings.Split(string(bytes), "\n")
  split := strings.Split(lines[0], " ")

  temp, _ := strconv.ParseInt(split[0], 10, 64)
  start := &node{temp, nil}
  current := start
  for i:=1;i<len(split);i++{
  fmt.Println(i, current)
    temp, _ = strconv.ParseInt(split[i], 10, 64)
    next := &node{temp, nil}
    current.next = next
    current = next
  }

  count := len(split)
  fmt.Println("")


  for i:=0;i<75;i++{
    fmt.Println(i, count)
    current = start
    for current != nil {
      val := current.val
      converted := strconv.FormatInt(val, 10)
      if val == 0 {
        current.val = 1
      } else if len(converted) & 1 == 0 {
        firstNum, _ := strconv.ParseInt(converted[:len(converted)/2], 10, 64)
        secondNum, _ := strconv.ParseInt(converted[len(converted)/2:], 10, 64)
        // fmt.Println(converted, firstNum, secondNum)
        new := &node{secondNum, nil}
        current.val = firstNum
        next := current.next
        if next == nil {
          current.next = new
          current = new
        } else {
          new.next = next
          current.next = new
          current = new
        }
        count++
      } else {
        current.val = val*2024
      }
      current = current.next
    }
  }
  //  current = start
  // for current != nil {
  //   fmt.Print(current.val, " ")
  //   current = current.next
  // }

  // current = start
  // for current != nil{
  //   count++
  //   current = current.next
  // }
  fmt.Println(count)
}

type node struct {
  val int64
  next *node
}
