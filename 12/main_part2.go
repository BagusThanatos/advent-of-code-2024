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
      area, sides := recursive(i,j,lines[i][j])
      if sides>0 {
        fmt.Println(string(lines[i][j]), i, j, area, sides)
      }
      sum += (area* sides)
    }
  }
  fmt.Println(sum)
}
var data [][]int64
func recursive(i, j int, c byte) (int64, int64) {
  // Note: number of sides is the same as the number of corners
  if (i<0 || j<0 || i>=len(lines) || j>=len(lines[i]) || lines[i][j] != c){
    return 0, 0
  }
  if data[i][j] == -1 {
    return 0, 0
  }
  if lines[i][j] != c {
    return 0, 0
  }
  data[i][j] = -1
  sides := int64(0)
  if (i-1<0 && j-1<0) ||
     (i-1<0&&j-1>=0&&lines[i][j-1]!=c) ||
     (j-1<0&&i-1>=0&&lines[i-1][j]!=c) ||
     (i-1>=0&&j-1>=0&&lines[i-1][j]!=c && lines[i][j-1]!=c){
    sides++
    fmt.Println("xax1", string(c), i, j)
  }
  if (i+1>=len(lines) && j-1<0) ||
    (i+1>=len(lines)&&j-1>=0&&lines[i][j-1]!=c) ||
    (i+1<len(lines) && j-1<0&&lines[i+1][j]!=c) ||
    (i+1<len(lines) && j-1>=0 && lines[i+1][j]!=c && lines[i][j-1]!=c){
    sides++
    fmt.Println("xax2", string(c), i, j)
  }
  if (i+1>=len(lines) && j+1>=len(lines[i])) ||
     (j+1<len(lines[i])&&i+1>=len(lines)&&lines[i][j+1]!=c) ||
     (i+1<len(lines) && j+1>=len(lines)&&lines[i+1][j]!=c) ||
     (j+1<len(lines[i]) && i+1<len(lines) && lines[i+1][j]!=c && lines[i][j+1]!=c){
    sides++
    fmt.Println("xax3", string(c), i, j)
  }
  if (i-1<0 && j+1>=len(lines[i])) ||
     (j+1<len(lines[i]) && i-1<0&&lines[i][j+1]!=c) ||
     (i-1>=0 && j+1>=len(lines[i])&&lines[i-1][j]!=c) ||
     (j+1<len(lines[i]) && i-1>=0 && lines[i-1][j]!=c && lines[i][j+1]!=c){
    sides++
    fmt.Println("xax4", string(c), i, j)
  }
  if (i-1>=0&&j-1>=0 && lines[i-1][j]==c && lines[i][j-1]==c && lines[i-1][j-1] != c){
    sides++
    fmt.Println("xax5", string(c), i, j)
  }
  if (i+1<len(lines)&&j-1>=0 && lines[i+1][j]==c && lines[i][j-1]==c && lines[i+1][j-1] != c){
    sides++
    fmt.Println("xax6", string(c), i, j)
  }
  if (i-1>=0&&j+1<len(lines[i]) && lines[i-1][j]==c && lines[i][j+1]==c && lines[i-1][j+1] != c){
    sides++
    fmt.Println("xax7", string(c), i, j)
  }
  if (i+1<len(lines)&&j+1<len(lines[i]) && lines[i+1][j]==c && lines[i][j+1]==c && lines[i+1][j+1] != c){
    sides++
    fmt.Println("xax8", string(c), i, j)
  }

  area := int64(1)
  deltaA, deltaP := recursive(i-1, j, c)
  area+=deltaA
  sides+=deltaP
  deltaA, deltaP = recursive(i+1, j, c)
  area+=deltaA
  sides+=deltaP
  deltaA, deltaP = recursive(i, j-1, c)
  area+=deltaA
  sides+=deltaP
  deltaA, deltaP = recursive(i, j+1, c)
  area+=deltaA
  sides+=deltaP
  fmt.Println("here", string(c), area, sides)
  return area, sides
}