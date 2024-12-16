package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "math"
)

var lines []string
var data [][]int64
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
  for i,_ := range lines {
    data[i] = make([]int64, len(lines[i]))
    for j, _ := range lines[i] {
      data[i][j] = int64(2*math.MaxInt32)
    }
  }

  // look for starting position
  y := 1
  x :=1
  for ;y<len(lines)-1;y++{
    x = 1
    for ;x<len(lines[y])-1;x++{
      if lines[y][x] == 'S' {
        goto found
      }
    }
  }
  found:
  // The idea is to find one best path first, and then look for all possible best path using the score as threshold
  recursiveFindFirst(RIGHT, y, x, 0)
  fmt.Println(string(lines[1][len(lines[1])-2]), data[1][len(lines[1])-2])
  recursiveFindAll(RIGHT, y, x, 0, []node{node{y, x}}, data[1][len(lines[1])-2])

  result := make(map[node]struct{})
  for i, v:= range data {
    fmt.Println(v)
    for j:=0;j<len(v);j++{
      if lines[i][j] == '#' {
        v[j] =2
      } else {
        v[j] = 0
      }
    }
  }
  for _, v := range bestList {
    result[v] = struct{}{}
    data[v.y][v.x] = 1 // let's mark it
  }
  for _, v:= range data {
    fmt.Println(v)
  }
  fmt.Println(len(result)+1)  // + 1 for the E node

}
var (
  UP = 0
  DOWN = 1
  LEFT = 2
  RIGHT = 3
)
type node struct{
  y, x int
}
var bestList []node
func recursiveFindFirst(mode, y, x int, current int64) {
  if y<0 || x<0 || y>=len(lines) || x>=len(lines[y]) || lines[y][x] == '#' || data[y][x] < current {
    return
  }
  if current < data[y][x] {
    data[y][x] = current
  }
  if lines[y][x] == 'E' {
    return
  }
  if mode == UP {
    recursiveFindFirst(mode, y-1, x, current + 1)
    recursiveFindFirst(RIGHT, y, x+1, 1001 + current)
    recursiveFindFirst(LEFT, y, x-1, 1001+ current)
  }
  if mode == DOWN {
    recursiveFindFirst(mode, y+1, x, current + 1)
    recursiveFindFirst(RIGHT, y, x+1, 1001 + current)
    recursiveFindFirst(LEFT, y, x-1, 1001+ current)
  }
  if mode == LEFT {
    recursiveFindFirst(mode, y, x-1, current + 1)
    recursiveFindFirst(UP, y-1, x, 1001 + current)
    recursiveFindFirst(DOWN, y+1, x, 1001+ current)
  }
  if mode == RIGHT {
    recursiveFindFirst(mode, y, x+1, current + 1)
    recursiveFindFirst(UP, y-1, x, 1001 + current)
    recursiveFindFirst(DOWN, y+1, x, 1001+ current)
  }
}
func recursiveFindAll(mode, y, x int, current int64, list []node, max int64) {
  if y<0 || x<0 || y>=len(lines) || x>=len(lines[y]) || lines[y][x] == '#' || current > max {
    return
  }
  if current == data[y][x] && lines[y][x] == 'E' {
    bestList = append(bestList, list...)
    return
  }
  list = append(list, node{y, x})
  if mode == UP {
    recursiveFindAll(mode, y-1, x, current + 1, list, max)
    recursiveFindAll(RIGHT, y, x+1, 1001 + current, list, max)
    recursiveFindAll(LEFT, y, x-1, 1001+ current, list, max)
  }
  if mode == DOWN {
    recursiveFindAll(mode, y+1, x, current + 1, list, max)
    recursiveFindAll(RIGHT, y, x+1, 1001 + current, list, max)
    recursiveFindAll(LEFT, y, x-1, 1001+ current, list, max)
  }
  if mode == LEFT {
    recursiveFindAll(mode, y, x-1, current + 1, list, max)
    recursiveFindAll(UP, y-1, x, 1001 + current, list, max)
    recursiveFindAll(DOWN, y+1, x, 1001+ current, list, max)
  }
  if mode == RIGHT {
    recursiveFindAll(mode, y, x+1, current + 1, list, max)
    recursiveFindAll(UP, y-1, x, 1001 + current, list, max)
    recursiveFindAll(DOWN, y+1, x, 1001+ current, list, max)
  }
}