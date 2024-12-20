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
  //sort by length
  // for i:=0;i<len(stripes);i++{
  //   for ii:=1;ii<len(stripes);ii++{
  //     if len(stripes[i]) < len(stripes[ii]) {
  //       temp := stripes[i]
  //       stripes[i] = stripes[ii]
  //       stripes[ii] = temp
  //     }
  //   }
  // }
  // build node
  start := &node{0, map[byte]*node{}}
  for _, stripe := range stripes {
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

/*
 Note: The idea is to "
 */

type node struct{
  c byte
  next map[byte]*node
}
var bestList []node
func recursiveNode(towel string, index int, currentNode *node) bool {
  if currentNode == nil {
    // fmt.Println("NULL NODE", towel, index)
    return false
  }
  if index >= len(towel) {
    // fmt.Println(towel, index, towel[index], string(towel[index]), string(currentNode.c), currentNode)
    // fmt.Println(towel, len(towel), index,  string(currentNode.c), currentNode)
    // return currentNode.next[0] != nil && (currentNode.c == towel[index] || currentNode.c == 0)
    // return currentNode.next[0] != nil || currentNode.c == 0//&& (currentNode.c == towel[index] || currentNode.c == 0)
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
func recursive(towel string, index int) bool {
  if index >= len(towel) {
    return true
  }
  result := false
  for i:=0;i<len(stripes);i++{
    stripe := stripes[i]
    if index+len(stripe) <= len(towel) && towel[index:index+len(stripe)] == stripe {
      // fmt.Println(index, stripe, towel[index:index+len(stripe)], stripe)
      result = result || recursive(towel, index + len(stripe))
      // fmt.Println(i, index, stripe, stripe)
    }
  }
  return result
}