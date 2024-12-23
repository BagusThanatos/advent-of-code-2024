package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  // "strconv"
  "math"
)

var lines []string
var data [][]int
var (
  SizeX = 141
  SizeY = 141
  wallCount = 1024
  MaxCheat = 2
  exitX = 1
  exitY = 1
)
func main() {
  input := "input_coba.txt"
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
  fmt.Println(lines)
  SizeY = len(lines)
  SizeX = len(lines[0])

  data = make([][]int, SizeY)
  for i := 0;i< SizeY; i++ {
    data[i] = make([]int, SizeX)
    for j:= 0;j<SizeX;j++{
      data[i][j]=math.MaxInt32
    }
  }
  // look for starting position
  y := 1
  x := 1
  for ;y<len(lines);y++{
    x =1
    for ;x<len(lines[y]);x++{
      if lines[y][x] == 'S'{
        goto findExit
      }
    }
  }
  findExit:
  // find exit
  for ;exitY<len(lines);exitY++{
    exitX = 1
    for ;exitX<len(lines[exitY]);exitX++{
      if lines[exitY][exitX] == 'E'{
        goto found
      }
    }
  }
  found:
  fmt.Println(y, x, exitY, exitX)
  recursiveFirst(y, x, 0)

  // Now we just need to print value in the location of 'E'
  for _, v:=range data {
    fmt.Println(v)
  }

  // lets cheat
  fastestWithoutCheating = data[exitY][exitX]
  // reset data array
  for i := 0;i< SizeY; i++ {
    for j:= 0;j<SizeX;j++{
      data[i][j]=math.MaxInt32
    }
  }
  fmt.Println(fastestWithoutCheating)
  recursive(y, x, 0, MaxCheat)
  fmt.Println(savesAtleast100)
  // for _, v:=range data {
  //   fmt.Println(v)
  // }

}
func recursiveFirst(y, x int, current int){
  if y<0 || x<0 || y>=SizeY || x>=SizeX || current >= data[y][x] {
    return
  }
  if lines[y][x] == '#'{
      return
  }
  if current < data[y][x] {
    data[y][x] = current
  }
  if y==exitY && x==exitX {
    return
  }
  recursiveFirst(y-1, x, current + 1)
  recursiveFirst(y, x+1, current + 1)
  recursiveFirst(y, x-1, current + 1)
  recursiveFirst(y+1, x, current + 1)
}
var savesAtleast100 = 0
var fastestWithoutCheating int
func recursive(y, x int, current int, cheat int){
  if y<0 || x<0 || y>=SizeY || x>=SizeX || current >= data[y][x] + 999{
    return
  }
  if lines[y][x] == '#'{
    if cheat > 0{
      cheat--
    } else {
      return
    }
  }
  if cheat==1 {
    cheat = 0
  }
  if current < data[y][x] {
    data[y][x] = current
  }
  if y==exitY && x==exitX {
    if current <= fastestWithoutCheating - 1{
      savesAtleast100++
    }
    return
  }
  recursive(y-1, x, current + 1, cheat)
  recursive(y, x+1, current + 1, cheat)
  recursive(y, x-1, current + 1, cheat)
  recursive(y+1, x, current + 1, cheat)
}