package main


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
  // lines = lines[:len(lines)-1]
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
  mode := '^'
  x := startX
  y := startY
  SIZE := 140
  walled := make([][]bool, SIZE)
  for i:=0;i<SIZE;i++{
    walled[i] = make([]bool, SIZE)
  }

  lines[y][x] = 'X'
  countX := 1

  for x>=0 && y >=0 && x<len(lines[0]) && y<len(lines) {
    fmt.Println(y, x)
    switch mode {
    case '^':
      if y-1>=0 {
        walled[y-1][x] = walled[y-1][x] || checkIfLoop(y-1, x, x, y, '>')
      }
    case '>':
      if x+1 < len(lines[0]) {
        walled[y][x+1] = walled[y][x+1] || checkIfLoop(y, x+1, x, y, 'v')
      }
    case 'v':
      if y+1 < len(lines) {
        walled[y+1][x] = walled[y+1][x] || checkIfLoop(y+1, x, x, y, '<')
      }
    case '<':
      // if y-1>=0 && lines[y-1][x] == 'X' {
      //   count++ // we can put an obstacle on the left side
      // }
      if x-1 >=0 {
        walled[y][x-1] = walled[y][x-1] || checkIfLoop(y, x-1, x, y, '^')
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
    if lines[y][x] != 'X' {
      lines[y][x] = 'X'
      countX++
    }
    if mode == '^'{
      if(y-1<0) {
        break
      }
      if lines[y-1][x] == '#'{
        mode = '>'
      }
    } else if mode == '>'{
      if x+1>=len(lines[0]){
        break
      }
      if lines[y][x+1] == '#'{
        mode = 'v'
      }
    } else if mode == 'v'{
      if y+1 >=len(lines) {
        break
      }
      if lines[y+1][x] == '#'{
        mode = '<'
      }
    }else if mode == '<'{
      if(x-1<0) {
        break
      }
      if lines[y][x-1] == '#'{
        mode = '^'
      }
    }
  }
  count := 0
  for i,v := range walled {
    if i<len(lines){
      // fmt.Println(v[:len(lines)])
    }
    for _, w := range v{
      if w {
        count++
      }
    }
  }
  fmt.Println(countX)
  fmt.Println(count)
}

func checkIfLoop(wallY, wallX, sx, sy int, mode rune) bool {
    // if wallX<0 || wallY<0 || wallX >= len(lines[0]) || wallY >= len(lines) {
    //   return false
    // }
    if lines[wallY][wallX] == '#' {
      return false
    }
    x:= sx
    y:= sy
    walked := 0
    for x>=0 && y >=0 && x<len(lines[0]) && y<len(lines) {
      // fmt.Println(y, x)
      if  (walked >= 100000){ // max walked, a sign of looping
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