package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "sort"
)

// NOte: Second part of day 23 is this: https://en.wikipedia.org/wiki/Clique_problem

var lines []string
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

  // from aa to zz
  connection := make([][]int, 26*26)
  for i := range connection {
    connection[i] = make([]int, 26*26)
  }
  for i:=0;i<len(lines);i++{
    nodes := strings.Split(lines[i], "-")
    firstNode := int(nodes[0][0]-'a')*26 + int(nodes[0][1]-'a')
    secondNode := int(nodes[1][0]-'a')*26 + int(nodes[1][1]-'a')
    connection[firstNode][secondNode] = 1
    connection[secondNode][firstNode] = 1
    fmt.Println(lines[i], nodes, firstNode, secondNode, string(nodes[1][0]), nodes[1][0], nodes[1][0] - 'a', int(nodes[1][0] - 'a')*26)
  }

  cliques := make([]int, 0)
  for i:=0;i<26*26;i++{
    tempCliques := []int{i}
    for secondNode := 0;secondNode<26*26;secondNode++{
      valid := true
      for _, firstNode := range tempCliques {
        if connection[firstNode][secondNode] != 1 {
          valid = false
          break
        }
      }
      if valid {
        tempCliques = append(tempCliques, secondNode)
      }
    }
    if len(tempCliques) > len(cliques) {
      cliques = tempCliques
    }
  }
  fmt.Println(len(cliques))
  sort.Ints(cliques)
  for _, v:= range cliques {
    fmt.Print(string('a'+v/26), string('a'+v%26), ",")
  }
}