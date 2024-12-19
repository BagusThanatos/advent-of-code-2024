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
  recursiveFindFirst(RIGHT, y, x, 0, []node{node{y, x}})
  fmt.Println(string(lines[1][len(lines[1])-2]), data[1][len(lines[1])-2])

  result := make(map[node]struct{})
  for i, v:= range data {
    fmt.Println(v)
    for j:=0;j<len(v);j++{
      if lines[i][j] == '#' {
        v[j] = 9
      } else {
        v[j] = 8
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
  fmt.Println(bestList)
  fmt.Println(result)
  fmt.Println(len(result)+1)  // + 1 for the E node

}
var (
  UP = 0
  DOWN = 1
  LEFT = 2
  RIGHT = 3
)

/*
 Note: The idea is to "look for all best times at the same time"
 1. Search the possibly more expensive path first: go for turns before going forward
 2. Handle possible cases when a node can be accessed from another path that's possibly more expensive at first, hence the 1001 addition below.
 it accounts for additional turn and a move
 */

type node struct{
  y, x int
}
var bestList []node
func recursiveFindFirst(mode, y, x int, current int64, list []node) {
  if y<0 || x<0 || y>=len(lines) || x>=len(lines[y]) || lines[y][x] == '#' || current > data[y][x] +1001 {
    return
  }
  if current < data[y][x] && lines[y][x] != 'E'{
    data[y][x] = current
  }
  if lines[y][x] == 'E' {
    if current < data[y][x]{
      // fmt.Println("replacing bestList with list", current, data[y][x], bestList, list)
      bestList = list
      data[y][x] = current
    } else if current == data[y][x] {
      // fmt.Println("appending bestlist with list", current, data[y][x], bestList, list)
      bestList = append(bestList, list...)
    }
    return
  }
  list = append(list, node{y, x})
  if mode == UP {
    recursiveFindFirst(RIGHT, y, x+1, 1001 + current, list)
    recursiveFindFirst(LEFT, y, x-1, 1001+ current, list)
    recursiveFindFirst(mode, y-1, x, current + 1, list)
  }
  if mode == DOWN {
    recursiveFindFirst(RIGHT, y, x+1, 1001 + current, list)
    recursiveFindFirst(LEFT, y, x-1, 1001+ current, list)
    recursiveFindFirst(mode, y+1, x, current + 1, list)
  }
  if mode == LEFT {
    recursiveFindFirst(UP, y-1, x, 1001 + current, list)
    recursiveFindFirst(DOWN, y+1, x, 1001+ current, list)
    recursiveFindFirst(mode, y, x-1, current + 1, list)
  }
  if mode == RIGHT {
    recursiveFindFirst(UP, y-1, x, 1001 + current, list)
    recursiveFindFirst(DOWN, y+1, x, 1001+ current, list)
    recursiveFindFirst(mode, y, x+1, current + 1, list)
  }
}