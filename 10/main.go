package main

import (
  "os"
  "fmt"
  "io"
  "strings"
)

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
        sum += len(t)
        fmt.Println(i, j, t)
      }
    }
  }
  fmt.Println(sum)
}

type coordinate struct {
  i int
  j int
}

func trail(i, j int, prev byte) map[coordinate]struct{} {
  if i<0 || j<0 || i>=len(lines) || j>= len(lines[i]) || lines[i][j] != prev + 1 {
    return nil
  }
  if lines[i][j] == '9' {
  fmt.Println(string(lines[i][j]))
    return map[coordinate]struct{}{coordinate{i,j}: {}}
  }
  adjacents := []map[coordinate]struct{}{trail(i+1, j, lines[i][j]), trail(i, j+1, lines[i][j]), trail(i-1, j, lines[i][j]),  trail(i, j-1, lines[i][j])}
  result := make(map[coordinate]struct{})
  for _, v := range adjacents {
    if v != nil {
      for key, val := range v{
        result[key] = val
      }
    }
  }
  return result
}
