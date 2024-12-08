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
  result := make([][]int64, len(lines))
  for i:=0;i<len(result);i++{
    result[i] = make([]int64, len(lines[0]))
  }
  for i:=0;i<len(lines);i++{
    for j:=0;j<len(lines[i]);j++{
      currentNode := lines[i][j]
      if currentNode == '.' {
        continue
      }
      // look for corresponding nodes
      for y:=0;y<len(lines);y++{
        for x:=0;x<len(lines);x++{
          if y==i && x==j {
            continue
          }
          if lines[y][x] == currentNode {
            // due to harmonics, if paired, each of them become antinodes
            result[i][j] = 1
            result[y][x] = 1
            newY := y + (y-i)
            newX := x + (x-j)
            // now we loop until out of bounds
            for newY >=0 && newY < len(lines) && newX >=0 && newX < len(lines[0]) {
              result[newY][newX] = 1
              newY += (y-i)
              newX += (x-j)
            }
          }
        }
      }
    }
  }

  sum := int64(0)
  for i:=0;i<len(result);i++{
      fmt.Println(result[i])
    for j:=0;j<len(result[i]);j++{
      sum+=result[i][j]
    }
  }
  fmt.Println(sum)
}
