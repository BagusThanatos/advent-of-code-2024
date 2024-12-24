package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "sort"
  // "math"
)

var lines []string
var m map[string]node
func main() {
  input := "input.txt"
  file, err := os.Open(input)
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

  m = make(map[string]node)
  index := 0
  for lines[index] != "" {
    strs := strings.Split(lines[index], ": ")
    value := 0
    if strs[1] == "1" {
      value = 1
    }
    m[strs[0]] = node{value, "", "", ""}
    index++
  }
  start:= index + 1

  index = start
  for index<len(lines) {
    strs := strings.Split(lines[index], " -> ")
    operation := strings.Split(strs[0], " ")
    m[strs[1]] = node{-1, operation[0], operation[1], operation[2]}
    index++
  }

  keys := make([]string, 0)
  for k := range m {
    if k[0] == 'z' {
      keys = append(keys, k)
    }
  }
  sort.Strings(keys)
  output := int64(0)
  // z00 is the rightmost bit, while z45 is the leftmost
  for k := len(keys) -1; k>=0;k-- {
    output = output << 1
    value := recursive(m[keys[k]])
    fmt.Println(k, keys[k], value)
    output = output | int64(value)
  }

  // fmt.Println(m)
  fmt.Println(output)
}
type node struct {
  value int
  parent1, operation, parent2 string
}
func recursive(current node) int{
  if current.value > -1 {
    return current.value
  }
  value := -1
  if current.operation == "AND" {
    if recursive(m[current.parent1]) == 1 &&  recursive(m[current.parent2]) == 1 {
      value = 1
    } else {
      value = 0
    }
  } else if current.operation == "OR" {
    if recursive(m[current.parent1]) == 1 || recursive(m[current.parent2]) == 1 {
      value = 1
    } else {
      value = 0
    }
  } else {
    if recursive(m[current.parent1]) != recursive(m[current.parent2]){
      value = 1
    } else {
      value = 0
    }
  }
  current.value = value
  return value
}