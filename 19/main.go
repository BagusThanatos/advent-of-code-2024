package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  // "math"
)

var lines []string
var stripes []string
func main() {
  // file, err := os.Open("input.txt")
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
  lines = lines[:len(lines)-1]

  stripes = strings.Split(lines[0], ", ")
  // NOTE: the idea here is to reduce the number of possible stripes so the problem has much smaller space
  //sort by length from shorter
  for i:=0;i<len(stripes);i++{
    for ii:=i+1;ii<len(stripes);ii++{
      if len(stripes[i]) > len(stripes[ii]) {
        temp := stripes[i]
        stripes[i] = stripes[ii]
        stripes[ii] = temp
      }
    }
  }
  // build tree
  start := &node{0, map[byte]*node{}}
  for _, stripe := range stripes {
    // only add to tree if it cant be build by existing tree
    if recursiveNode(stripe, 0, start) {
      fmt.Println("skipped", stripe)
      continue
    }
    current := start
    for i := range stripe {
      c := stripe[i]
      if current.next[c] == nil {
        current.next[c] = &node{c, map[byte]*node{}}
      }
      current = current.next[c]
    }
    current.next[0] = start
  }
  delete(start.next, 0)

  count := int64(0)
  for i:=2;i<len(lines);i++{
    result := recursiveNode(lines[i], 0, start)
    fmt.Println(lines[i], result)
    if result {
      count++
    }
  }
  fmt.Println(count)
}
var (
  UP = 0
  DOWN = 1
  LEFT = 2
  RIGHT = 3
)


type node struct{
  c byte
  next map[byte]*node
}
var bestList []node
func recursiveNode(towel string, index int, currentNode *node) bool {
  if currentNode == nil {
    return false
  }
  if index >= len(towel) {
    return currentNode.next[0] != nil
  }
  next := currentNode.next[towel[index]]
  result := recursiveNode(towel, index+1, next)
  next = currentNode.next[0]
  if next != nil {
    result = result || recursiveNode(towel, index+1, next.next[towel[index]])
  }
  return result
}