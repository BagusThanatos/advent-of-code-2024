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

  harddisk := make([]int64, 1024*1024) // 1MB of harddisk

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
  for ;int64(emptyBlockLoc)<currentHDIndex && harddisk[emptyBlockLoc] != EMPTYBLOCK ; emptyBlockLoc++ {}
  fmt.Println(currentHDIndex, harddisk[currentHDIndex])
  currentHDIndex--
  currentFile := harddisk[currentHDIndex]
  for currentHDIndex >=0 && int64(emptyBlockLoc) < currentHDIndex{
    // determine file length
    length := 0
    for harddisk[currentHDIndex] == currentFile{
      currentHDIndex--
      length++
    }
    // look for longest empty block that can fit from leftmost
    emptyBlockLoc := 0
    currentEmptyBlockLength := 0
    for emptyBlockLoc < currentHDIndex && currentEmptyBlockLength < length {
      if harddisk[emptyBlockLoc] == EMPTYBLOCK {
        currentEmptyBlockLength++
      } else {
        currentEmptyBlockLength = 0
      }
      emptyBlockLoc++
    }

    // move blocks
    if currentEmptyBlockLength == length {
      for i:=0;i<length;i++{
        harddisk[emptyBlockLoc-i] = harddisk[currentHDIndex+i]
        harddisk[currentHDIndex+i] = EMPTYBLOCK
      }
    }
  }

  sum := int64(0)
  for i := 0;i<len(harddisk)&&harddisk[i] != EMPTYBLOCK; i++{
    sum += int64(i) * harddisk[i]
  }
  fmt.Println(harddisk[:100])
  fmt.Println(sum)
}
