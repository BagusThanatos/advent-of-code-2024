package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
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
  input := lines[0]

  harddisk := make([]int64, 1024*1024) // 1KB of harddisk

  EMPTYBLOCK := int64(-1)
  currentFileID := int64(0)
  currentHDIndex := int64(0)

  // generate block rules first
  fmt.Println("generate block rules")
  for i, v := range input {
    value, _ := strconv.ParseInt(string(v), 10, 64)
    set := EMPTYBLOCK
    if i & 1 == 0 { // file indices
      set = currentFileID
      currentFileID++
    }
    until := currentHDIndex + value
    for currentHDIndex < until {
      harddisk[currentHDIndex] = set
      currentHDIndex++
    }
  }
  fmt.Println("move blocks")
  emptyBlockLoc := 0
  for ;int64(emptyBlockLoc)<currentHDIndex && harddisk[emptyBlockLoc] != EMPTYBLOCK ; emptyBlockLoc++ {}
  fmt.Println(currentHDIndex, harddisk[currentHDIndex])
  currentHDIndex--
  for currentHDIndex >=0 && int64(emptyBlockLoc) < currentHDIndex{
    harddisk[emptyBlockLoc] = harddisk[currentHDIndex]
    harddisk[currentHDIndex] = EMPTYBLOCK
    for ;currentHDIndex>=0 && harddisk[currentHDIndex] == EMPTYBLOCK; currentHDIndex--{}
    for ;emptyBlockLoc<len(harddisk) && harddisk[emptyBlockLoc] != EMPTYBLOCK ; emptyBlockLoc++ {}
    if emptyBlockLoc >= len(harddisk) {
      emptyBlockLoc = 0
      for ;int64(emptyBlockLoc) < currentHDIndex && emptyBlockLoc<len(harddisk) && harddisk[emptyBlockLoc] != EMPTYBLOCK ; emptyBlockLoc++ {}
    }
  }

  sum := int64(0)
  for i := 0;i<len(harddisk)&&harddisk[i] != EMPTYBLOCK; i++{
    sum += int64(i) * harddisk[i]
  }
  fmt.Println(sum)
}
