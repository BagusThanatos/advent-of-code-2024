package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
  "math"
)

var lines []string
var data [][]int
var walls [][]int
var (
  SizeX = 71
  SizeY = 71
  wallCount = 1024
)
func main() {
  input := "input.txt"
  if strings.Contains(input, "coba") {
    SizeY = 7
    SizeX = 7
    wallCount = 12
  }

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
  fmt.Println(len(lines))

  data = make([][]int, SizeY)
  walls = make([][]int, SizeY)
  for i := 0;i< SizeY; i++ {
    data[i] = make([]int, SizeX)
    for j:= 0;j<SizeX;j++{
      data[i][j]=math.MaxInt32
    }
    walls[i] = make([]int, SizeX)
  }

  // reset data
  data[SizeY-1][SizeX-1] = math.MaxInt32-1
  for data[SizeY-1][SizeX-1] != math.MaxInt32{
    for i := 0;i< SizeY; i++ {
      for j:= 0;j<SizeX;j++{
        data[i][j]=math.MaxInt32
      }
    }
    for i:=0;i<wallCount;i++{
      v := lines[i]
      coordinates := strings.Split(v, ",")
      x, _ := strconv.ParseInt(coordinates[0], 10, 32)
      y, _ := strconv.ParseInt(coordinates[1], 10, 32)
      walls[y][x] = 1
    }
    // look for starting position
    y := 0
    x := 0
    recursive(y, x, 0)
    wallCount++
  }

  // Now we just need to print value in the location of 'E'
  for _, v:=range data {
    fmt.Println(v)
  }
  fmt.Println(wallCount, lines[wallCount-2], data[SizeY-1][SizeX-1])

}
var (
  UP = 0
  DOWN = 1
  LEFT = 2
  RIGHT = 3
)
func recursive(y, x int, current int){
  if y<0 || x<0 || y>=SizeY || x>=SizeX || walls[y][x] == 1 || current >= data[y][x] {
    return
  }
  if current < data[y][x] {
    data[y][x] = current
  }
  if y==SizeY-1 && x==SizeX-1 {
    return
  }
  recursive(y-1, x, current + 1)
  recursive(y, x+1, current + 1)
  recursive(y, x-1, current + 1)
  recursive(y+1, x, current + 1)
}