package main

import (
  "os"
  "fmt"
  "io"
  // "math"
  "strings"
  "strconv"
)

// Funny story: I stumbled into the solution for the second part first lmao
var lines []string
var (
  PERIOD = int64(100)
  // For example input
  // MAXX = int64(11)
  // MAXY = int64(7)


  // for actual input
  MAXX = int64(101)
  MAXY = int64(103)
  QUADRANTX = MAXX/2
  QUADRANTY = MAXY/2
)
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
  quadrants := [][]int64{[]int64{0, 0}, []int64{0, 0}}
  // data := make([][]int, MAXY)
  // for i:=0;i<int(MAXY);i++{
  //   data[i] = make([]int, MAXX)
  // }
  for i := 0;i<len(lines);i++{
    posX, posY, vX, vY := parseRobot(lines[i])
    finalX := (posX + (PERIOD * vX)) % MAXX
    finalY := (posY + (PERIOD * vY)) % MAXY
    if finalX < 0 {
      finalX = MAXX + finalX
    }
    if finalY < 0 {
      finalY = MAXY + finalY
    }
    // fmt.Println(posX, posY, vX, vY)
    // fmt.Println(finalX, finalY, posX + (PERIOD * vX), )
    if finalX == QUADRANTX || finalY == QUADRANTY {
      continue
    }
    // data[int(finalY)][int(finalX)] += 1
    quadrants[finalX/(QUADRANTX+1)][finalY/(QUADRANTY+1)] += 1
    // fmt.Println(quadrants)
  }
  sum := int64(1)
  for _, v := range quadrants {
    for _, x := range v {
      sum *= x
    }
  }
  // fmt.Println(quadrants)
  // for _, v := range data {
  //   fmt.Println(v)
  // }
  fmt.Println(sum)

}
func parseRobot(s string) (int64, int64, int64, int64) {
  parsed := strings.Split(s, " ")
  pos := strings.Split(strings.Split(parsed[0], "=")[1], ",")
  posX, _ := strconv.ParseInt(pos[0], 10, 64)
  posY, _ := strconv.ParseInt(pos[1], 10, 64)
  v := strings.Split(strings.Split(parsed[1], "=")[1], ",")
  vX, _ := strconv.ParseInt(v[0], 10, 64)
  vY, _ := strconv.ParseInt(v[1], 10, 64)
  return posX, posY, vX, vY
}
func abs(i int64) int64{
  if i<0{
    return -i
  }
  return i
}
