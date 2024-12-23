package main

import (
  "os"
  "fmt"
  "io"
  "strings"
)

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
  // for all node that starts with a 't'
  set := make(map[node]struct{})
  for i:=0;i<=26;i++{
    firstNode := int('t'-'a') * 26 + i
    for secondNode := 0;secondNode<26*26;secondNode++{
      if connection[firstNode][secondNode]!=1 {
        continue
      }
      for thirdNode := 0;thirdNode<26*26;thirdNode++{
        if connection[secondNode][thirdNode] != 1 || connection[thirdNode][firstNode] != 1 {
          continue
        }
        nums := []int{firstNode, secondNode, thirdNode}
        for y:=0;y<len(nums);y++{
          for x:=y+1;x<len(nums);x++{
            if nums[x] < nums[y] {
              temp := nums[x]
              nums[x] = nums[y]
              nums[y] = temp
            }
          }
        }
        set[node{nums[0],nums[1], nums[2]}] = struct{}{}
      }
    }
  }
  count := 0
  for _ = range set {
    count++
  }
  fmt.Println(count)
}
type node struct{
  a,b,c int
}
