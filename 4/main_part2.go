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
    if j-1>=0 && j+1 < len(lines[i]) && i+1<len(lines) && i-1>=0 && lines[i][j] == 'A' && lines[i+1][j+1] == 'S' && lines[i-1][j-1] == 'M' && lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S'{
      count++
    } else if j-1>=0 && j+1 < len(lines[i]) && i+1<len(lines) && i-1>=0 && lines[i][j] == 'A' && lines[i+1][j+1] == 'S' && lines[i-1][j-1] == 'M' && lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M'{
      count++
    } else if j-1>=0 && j+1 < len(lines[i]) && i+1<len(lines) && i-1>=0 && lines[i][j] == 'A' && lines[i+1][j+1] == 'M' && lines[i-1][j-1] == 'S' && lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M'{
      count++
    } else if j-1>=0 && j+1 < len(lines[i]) && i+1<len(lines) && i-1>=0 && lines[i][j] == 'A' && lines[i+1][j+1] == 'M' && lines[i-1][j-1] == 'S' && lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S'{
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