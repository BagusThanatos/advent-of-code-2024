package main

import (
  "os"
  "fmt"
  "io"
  "strings"
)

var lines []string
var stripes []string
var stripeBase [][]string
func main() {
  file, err := os.Open("input.txt")
  // file, err := os.Open("input_coba.txt")
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
      if len(stripes[i]) >len(stripes[ii]) {
        temp := stripes[i]
        stripes[i] = stripes[ii]
        stripes[ii] = temp
      }
    }
  }
  // fmt.Println(stripes)
  // build node
  start := &node{"", map[byte]*node{}}
  for _, stripe := range stripes {
    // only add to tree if it cant be build by existing tree
    if ok, _ := recursiveNode(stripe, 0, start, "", []string{}); ok{
      // fmt.Println("skipped", stripe)
      continue
    }
    fmt.Println("added to tree:", stripe)
    current := start
    for i := range stripe {
      c := stripe[i]
      if current.next[c] == nil {
        current.next[c] = &node{"", map[byte]*node{}}
      }
      current = current.next[c]
    }
    current.next[0] = start
    fmt.Println(current.stripe, stripe)
    current.stripe = stripe
  }
  delete(start.next, 0)
  stripeBase = make([][]string, 0)
  for _, stripe := range stripes {
    _, m := recursiveNode(stripe, 0, start, "", []string{})
    // fmt.Println(stripe, m, result)
    stripeBase = append(stripeBase, m)
    // stripeBase[stripe] = m
  }
  count := int64(0)
  for i:=2;i<len(lines);i++{
    result, m := recursiveNode(lines[i], 0, start, "", []string{})
    // result := recursive(lines[i], 0)
    // fmt.Println(lines[i], result, m)
    if result {
      tempCount := recursivePossibleStripes(m, 0)
      fmt.Println("possible for", lines[i], m, tempCount)
      count+=int64(tempCount)
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

func recursivePossibleStripes(towel []string, index int) int {
  if index >= len(towel) {
    return 1
  }
  count := 0
  for i := range stripeBase {
    if index + len(stripeBase[i]) <= len(towel) && sliceEqual(towel[index:index+len(stripeBase[i])], stripeBase[i]) {
      count += recursivePossibleStripes(towel, index+len(stripeBase[i]))
    }
  }
  return count
}
func sliceEqual(a, b []string) bool {
  if len(a) != len(b){
    return false
  }
  for i:=0;i<len(a);i++{
    if a[i] != b[i]{
      return false
    }
  }
  return true
}
type node struct{
  stripe string
  next map[byte]*node
}
var bestList []node
func recursiveNode(towel string, index int, currentNode *node, currentStripe string, list []string) (bool, []string) {
  // fmt.Println(towel, index, currentNode)
  if currentNode == nil {
    // fmt.Println("NULL NODE", towel, index)
    return false, []string{}
  }
  if index >= len(towel) {
    // fmt.Println(towel, index, towel[index], string(towel[index]), string(currentNode.c), currentNode)
    // return currentNode.next[0] != nil && (currentNode.c == towel[index] || currentNode.c == 0)
    // return currentNode.next[0] != nil || currentNode.c == 0//&& (currentNode.c == towel[index] || currentNode.c == 0)
    if currentNode.next[0] == nil {
      return false, []string{}
    }
      // fmt.Println(towel, len(towel), index,  string(currentNode.stripe), currentNode, append(list, currentStripe))
    return true, append(list, currentStripe)
  }
  next := currentNode.next[towel[index]]
  result, m := recursiveNode(towel, index+1, next, currentStripe+string(towel[index]), list)
  if result {
    return result, m
  }
  next = currentNode.next[0]
  if next != nil {
    tempResult, tempM := recursiveNode(towel, index+1, next.next[towel[index]], string(towel[index]), append(list, currentStripe))
    // if !result {
      // fmt.Println("HERE", tempResult, tempM, m)
      // m = tempM
    // }
    // result = result || tempResult
    return tempResult, tempM
  }
  return false, []string{}
}