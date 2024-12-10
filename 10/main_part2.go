package main

import (
  "os"
  "fmt"
  "io"
  "strings"
)

// Funny story: I stumbled into the solution for the second part first lmao
var lines []string
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

  sum:=0
  // look for start
  for i:=0;i<len(lines);i++{
    for j:=0;j<len(lines[i]);j++{
      if lines[i][j]=='0'{
        t := trail(i, j, '0' -1 )
        sum += t
      }
    }
  }
  fmt.Println(sum)
}

type coordinate struct {
  i int
  j int
}

func trail(i, j int, prev byte) int {
  if i<0 || j<0 || i>=len(lines) || j>= len(lines[i]) || lines[i][j] != prev + 1 {
    return 0
  }
  if lines[i][j] == '9' {
    return 1
  }
  return trail(i+1, j, lines[i][j]) + trail(i, j+1, lines[i][j]) + trail(i-1, j, lines[i][j]) +  trail(i, j-1, lines[i][j])

}
