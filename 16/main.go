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
  recursive(RIGHT, y, x, 0)
  fmt.Println(string(lines[1][len(lines[1])-2]), data[1][len(lines[1])-2])
  // for _, v:= range data {
  //   fmt.Println(v)
  // }

}
var (
  UP = 0
  DOWN = 1
  LEFT = 2
  RIGHT = 3
)
func recursive(mode, y, x int, current int64) {
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
    recursive(mode, y-1, x, current + 1)
    recursive(RIGHT, y, x+1, 1001 + current)
    recursive(LEFT, y, x-1, 1001+ current)
  }
  if mode == DOWN {
    recursive(mode, y+1, x, current + 1)
    recursive(RIGHT, y, x+1, 1001 + current)
    recursive(LEFT, y, x-1, 1001+ current)
  }
  if mode == LEFT {
    recursive(mode, y, x-1, current + 1)
    recursive(UP, y-1, x, 1001 + current)
    recursive(DOWN, y+1, x, 1001+ current)
  }
  if mode == RIGHT {
    recursive(mode, y, x+1, current + 1)
    recursive(UP, y-1, x, 1001 + current)
    recursive(DOWN, y+1, x, 1001+ current)
  }
}