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
  currentHDIndex := 0

  // generate block rules first
  fmt.Println("generate block rules")
  for i, v := range input {
    value, _ := strconv.ParseInt(string(v), 10, 64)
    set := EMPTYBLOCK
    if i & 1 == 0 { // file indices
      set = currentFileID
      currentFileID++
    }
    until := currentHDIndex + int(value)
    for currentHDIndex < until {
      harddisk[currentHDIndex] = set
      currentHDIndex++
    }
  }
  fmt.Println(harddisk[:100])
  fmt.Println("move blocks")
  fmt.Println(currentHDIndex, harddisk[currentHDIndex])
  currentHDIndex--
  for currentHDIndex >=0{
    // determine file length
    // fmt.Println("determine file length")
    length := 0
    currentFile := harddisk[currentHDIndex]
    for currentHDIndex >= 0 && harddisk[currentHDIndex] == currentFile{
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
    // fmt.Println(currentFile, length, currentHDIndex, currentEmptyBlockLength, emptyBlockLoc, currentEmptyBlockLength == length)

    // move blocks
    if currentEmptyBlockLength == length {
      // fmt.Println("move blocks for ", currentFile)
      for i:=1;i<=length;i++{
        harddisk[emptyBlockLoc-i] = harddisk[currentHDIndex+i]
        harddisk[currentHDIndex+i] = EMPTYBLOCK
      }
      // fmt.Println(harddisk[:50])
    }
    for ;currentHDIndex>=0 && harddisk[currentHDIndex] == EMPTYBLOCK; currentHDIndex--{}
  }

  sum := int64(0)
  for i := 0;i<len(harddisk); i++{
    if harddisk[i] > 0 {
      sum += (int64(i) * harddisk[i])
    }
  }
  // try to validate result
  // valid := map[int64][]int{}
  // curr := harddisk[0]
  // valid[curr] = []int{0}
  // for i := 1;i<len(harddisk); i++{
  //   if harddisk[i] > 0 && harddisk[i] != curr {
  //     if v, ok := valid[curr]; ok {
  //       valid[curr] = append(v, i)
  //     } else {
  //       valid[curr] = []int{i}
  //     }
  //     curr = harddisk[i]
  //   }
    // if harddisk[i] == 0 {
    //   fmt.Println(i)
    // }
  // }
  // for k, v := range valid {
  //   if len(v) > 1 {
  //     fmt.Println("occur twice: ", k, v)
  //   }
  // }
  // fmt.Println(valid)
  fmt.Println(harddisk[:100])
  fmt.Println(sum)
}
