package main


import (
  "os"
  "fmt"
  "io"
  "strings"
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

  linesString := strings.Split(string(bytes), "\n")
  // lines = lines[:len(lines)-1]
  lines := make([][]byte, len(linesString)-1)
  for i:=0;i<len(lines);i++{
    lines[i] = []byte(linesString[i])
  }
  fmt.Println(lines[0][0] == '^')

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
  fmt.Println(startX, startY)
  mode := '^'
  x := startX
  y := startY
  count := 1
  lines[y][x] = 'X'
  for x>=0 && y >=0 && x<len(lines[0]) && y<len(lines) {
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
    if lines[y][x] != 'X'{
      count++
      lines[y][x] = 'X'
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
  fmt.Println(count)
}