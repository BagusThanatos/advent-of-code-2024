package main

/*
NOTE: This is the correct answer
The problem was to put an obstacle BEFORE the guard start patrolling
*/

import (
  "os"
  "fmt"
  "io"
  "strings"
)

var lines [][]byte
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

  linesString := strings.Split(string(bytes), "\n")
  lines = make([][]byte, len(linesString)-1)
  for i:=0;i<len(lines);i++{
    lines[i] = []byte(linesString[i])
  }

  // look for starting position
  startX := 0
  startY := 0
  for ;startX<len(lines[0]);startX++ {
    for startY=0;startY < len(lines);startY++{
      if lines[startY][startX] == '^' {
        goto walk
      }
    }
  }
  walk:
  // walk the thing
  fmt.Println(startX, startY, string(lines[startY][startX]))
  fmt.Println(len(lines), len(lines[0]))
  SIZE := len(lines)
  walled := make([][]bool, SIZE)
  for i:=0;i<SIZE;i++{
    walled[i] = make([]bool, SIZE)
  }

  x := 0
  for x<len(lines[0])  {
    y := 0
    for y<len(lines) {
      walled[y][x] = checkIfLoop(y, x, startX, startY, '^')
      y++
    }
    x++
  }
  count := 0
  walled[startY][startX] = false
  for i,v := range walled {
    if i<len(lines){
      fmt.Println(v[:len(lines)])
    }
    for _, w := range v{
      if w {
        count++
      }
    }
  }
  fmt.Println(count)
}

func checkIfLoop(wallY, wallX, sx, sy int, mode rune) bool {
    if lines[wallY][wallX] == '#' {
      return false
    }
    x:= sx
    y:= sy
    walked := 0
    for x>=0 && y >=0 && x<len(lines[0]) && y<len(lines) {
      // fmt.Println(y, x)
      if  (walked >= 900000){ // max walked, a sign of looping
        return true
      }
      walked++
      if mode == '^'{
        if(y-1<0) {
          break
        }
        if lines[y-1][x] == '#' || (y-1==wallY && x==wallX){
          mode = '>'
          continue
        }
      } else if mode == '>'{
        if x+1>=len(lines[0]){
          break
        }
        if lines[y][x+1] == '#'||(y==wallY && x+1==wallX){
          mode = 'v'
          continue
        }
      } else if mode == 'v'{
        if y+1 >=len(lines) {
          break
        }
        if lines[y+1][x] == '#'||(y+1==wallY && x==wallX){
          mode = '<'
          continue
        }
      }else if mode == '<'{
        if(x-1<0) {
          break
        }
        if lines[y][x-1] == '#'||(y==wallY && x-1==wallX){
          mode = '^'
          continue
        }
      }
      switch mode {
      case '^':
        y--
      case '>':
        x++
      case 'v':
        y++
      case '<':
        x--
      }
    }
    return false
  }