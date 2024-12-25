package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  // "sort"
  // "math"
)

var lines []string
func main() {
  input := "input.txt"
  file, err := os.Open(input)
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

  locks := make([][]int, 0)
  keys := make([][]int, 0)
  index := 0
  for index < len(lines) {
    item := lines[index:index+7]
    v := make([]int, 5)
    for x:=0;x<5;x++{
      count := 0
      for y:=0;y<7;y++{
        if item[y][x] == '#' {
          count++
        }
      }
      v[x] = count
    }
    if item[0] == "#####" {
      // a lock
      locks = append(locks, v)
    } else {
      // a key
      keys = append(keys, v)
    }
    index += 8
  }
  fmt.Println(locks[:4])
  fmt.Println(keys[:4])
  count := 0
  for _, lock := range locks {
    for _, key := range keys {
      valid := true
      for x:=0;valid && x<len(lock);x++{
        if lock[x] + key[x] > 7 {
          valid = false
        }
      }
      if valid{
        count ++
      }
    }
  }
  fmt.Println(count)
}
