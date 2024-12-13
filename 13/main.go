package main

import (
  "os"
  "fmt"
  "io"
  "math"
  "strings"
  "strconv"
)

// Funny story: I stumbled into the solution for the second part first lmao
var lines []string
var (
  COSTA = 3
  COSTB = 1
  MAXPRESSED = 100
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
  // lines = lines[:len(lines)-1]
  sum := int64(0)
  for i := 0;i<len(lines);i+=4{
    butA := lines[i]
    butB := lines[i+1]
    prize := lines[i+2]
    deltaXA, deltaYA := parseButton(butA)
    deltaXB, deltaYB := parseButton(butB)
    prizeX, prizeY := parseButton(prize)
    cost := math.MaxInt
    for a:=1;a<=MAXPRESSED;a++{
      for b:=1;b<=MAXPRESSED;b++{
        if a*deltaXA + b*deltaXB == prizeX && a*deltaYA + b*deltaYB == prizeY {
          tempCost := COSTA*a + COSTB * b
          if tempCost < cost {
            fmt.Println(a, b)
            cost = tempCost
          }
        }
      }
    }
    if cost < math.MaxInt {
      sum += int64(cost)
    }
  }
  fmt.Println(sum)

}
var data [][]int64
func parseButton(s string) (int, int) {
  parsed := strings.Split(strings.Split(s, ":")[1], ",")
  deltaX, _ := strconv.ParseInt(parsed[0][3:], 10, 32)
  deltaY, _ := strconv.ParseInt(parsed[1][3:], 10, 32)
  return int(deltaX), int(deltaY)
}