package main

import (
  "os"
  "fmt"
  "io"
  // "math"
  "strings"
  "strconv"
)

// Answer is: 8053, Found through actually looking into the printed results lol
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
  data := make([][]int, MAXY)
  for i:=0;i<int(MAXY);i++{
    data[i] = make([]int, MAXX)
  }
  // 2999 is too low
  elapsed := int64(3000)

  for elapsed < 10000 {
    for i := 0;i<len(lines);i++{
      posX, posY, vX, vY := parseRobot(lines[i])
      finalX := (posX + (elapsed * vX)) % MAXX
      finalY := (posY + (elapsed * vY)) % MAXY
      if finalX < 0 {
        finalX = MAXX + finalX
      }
      if finalY < 0 {
        finalY = MAXY + finalY
      }
      data[finalY][finalX] = 7
    }
    fmt.Println(elapsed)
    for _, v:= range data {
      for _, x := range v {
        fmt.Print(x)
      }
      fmt.Println()
    }
    for i:=0;i<int(MAXY);i++{
      data[i] = make([]int, MAXX)
    }
    elapsed++
  }
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