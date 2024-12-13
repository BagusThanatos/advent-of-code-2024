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
  COSTA = int64(3)
  COSTB = int64(1)
  MAXPRESSED = int64(100)
)

/*
The problem is basically asking us to solve a linear equation with two variables
CXa + DXb = X
CYa + DYb = Y

e.g
first equation
11a + 17b = 2413
second equation
67a + 16b = 6293

multiply first equation with 16
176a + 272b = 38608
multiply second equation with 17
1139a+ 272b = 106981

substraction
963a = 68373
*/
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
    deltaXA, deltaYA := parseButton(butA, 0)
    deltaXB, deltaYB := parseButton(butB, 0)
    prizeX, prizeY := parseButton(prize, 10000000000000)
    // prizeX, prizeY := parseButton(prize, 0)
    // odd numbers require one of the component to be BOTH(deltaX and any multipliers) odds so we can skipp if deltas are even
    if (prizeX&1!=0 && deltaXA &1==0&& deltaXB&1==0) || (prizeY&1!=0 && deltaYA &1==0&& deltaYB&1==0){
      continue
    }
    a := abs((prizeY*deltaXB-prizeX*deltaYB)/(deltaXA*deltaYB - deltaYA*deltaXB))
    b := (prizeX-deltaXA*a)/deltaXB
    // fmt.Println(deltaXA, deltaXB, prizeX)
    // fmt.Println(deltaYA, deltaYB, prizeY)
    // fmt.Println(a, b)
    if deltaXA*a + (deltaXB*b) != prizeX || (deltaYA*a) + (deltaYB*b) != prizeY{
      continue
    }
    cost := int64(a*COSTA+b*COSTB)
    fmt.Println(a, b, cost)
    if cost < math.MaxInt64{
      sum += cost
      fmt.Println(cost, i)
    }
  }
  fmt.Println(sum)

}
var data [][]int64
func parseButton(s string, add int64) (int64, int64) {
  parsed := strings.Split(strings.Split(s, ":")[1], ",")
  deltaX, _ := strconv.ParseInt(parsed[0][3:], 10, 64)
  deltaY, _ := strconv.ParseInt(parsed[1][3:], 10, 64)
  return add + deltaX, add + deltaY
}
func abs(i int64) int64{
  if i<0{
    return -i
  }
  return i
}
