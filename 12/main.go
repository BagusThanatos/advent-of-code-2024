package main

import (
  "os"
  "fmt"
  "io"
  "strings"
)

// Funny story: I stumbled into the solution for the second part first lmao
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
  lines = lines[:len(lines)-1]
  data = make([][]int64, len(lines))
  for i:=0;i<len(data);i++{
    data[i] = make([]int64, len(lines[i]))
  }
  fmt.Println(data, len(data), len(data[1]))

  sum := int64(0)
  for i:=0;i<len(lines);i++{
    for j:=0;j<len(lines[i]);j++{
      area, perimeter := recursive(i,j,lines[i][j])
      if perimeter>0 {
        fmt.Println(string(lines[i][j]), i, j, area, perimeter)
      }
      sum += (area* perimeter)
    }
  }
  fmt.Println(sum)
}
var data [][]int64
func recursive(i, j int, c byte) (int64, int64) {
  if (i<0 || j<0 || i>=len(lines) || j>=len(lines[i]) || lines[i][j] != c){
    return 0, 1
  }
  if data[i][j] == -1 {
    return 0, 0
  }
  if lines[i][j] != c {
    return 0, 0
  }
  data[i][j] = -1
  perimeter := int64(0)
  area := int64(1)
  deltaA, deltaP := recursive(i-1, j, c)
  area+=deltaA
  perimeter+=deltaP
  deltaA, deltaP = recursive(i+1, j, c)
  area+=deltaA
  perimeter+=deltaP
  deltaA, deltaP = recursive(i, j-1, c)
  area+=deltaA
  perimeter+=deltaP
  deltaA, deltaP = recursive(i, j+1, c)
  area+=deltaA
  perimeter+=deltaP
  return area, perimeter
}