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

  lines := strings.Split(string(bytes), "\n")
  lines = lines[:len(lines)-1]

  sum := int64(0)

  checkXMAS := func(i, j int) int64 {
    count := int64(0)
    if j+3 < len(lines[i]) && lines[i][j] == 'X' && lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
      count++
    }
    if j-3 >= 0 && lines[i][j] == 'X' && lines[i][j-1] == 'M' && lines[i][j-2] == 'A' && lines[i][j-3] == 'S' {
      count++
    }
    if j+3 < len(lines[i]) && i+3 < len(lines) && lines[i][j] == 'X' && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
      count++
    }
    if j-3 >= 0 && i-3 >= 0 && lines[i][j] == 'X' && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
      count++
    }
    if j+3 < len(lines[i]) && i-3 >= 0 && lines[i][j] == 'X' && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
      count++
    }
    // fmt.Println(i, j)
    if j-3 >= 0 && i+3 < len(lines) && lines[i][j] == 'X' && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
      count++
    }
    if i-3 >= 0 && lines[i][j] == 'X' && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
      count++
    }
    if i+3 < len(lines) && lines[i][j] == 'X' && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
      count++
    }

    return count
  }

  for i:=0;i<len(lines);i++ {
    for j:=0;j<len(lines[i]); j++ {
      sum += checkXMAS(i, j)
    }
  }
  fmt.Println(sum)
}