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

  harddisk := make([]int64, 94521) // Length was from it's maximum elements

  EMPTYBLOCK := int64(-1)
  currentFileID := int64(0)
  currentHDIndex := 0
  fmt.Println(len(lines[0]))
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
  PrintHDD(harddisk)
  fmt.Println("move blocks")
  currentHDIndex--
  fmt.Println(currentHDIndex, harddisk[currentHDIndex])
  for currentHDIndex >=0{
    // determine file length
    length := 0
    currentFile := harddisk[currentHDIndex]
    for currentHDIndex >= 0 && harddisk[currentHDIndex] == currentFile{
      currentHDIndex--
      length++
    }
    // look for longest empty block that can fit from leftmost
    emptyBlockLoc := 0
    currentEmptyBlockLength := 0
    // NOTE: was stuck here due to `< currentHDIndex` not covering cases such `-1 -1 -1 1 1 1`
    for emptyBlockLoc <= currentHDIndex && currentEmptyBlockLength < length {
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
      for i:=1;i<=length;i++{
        harddisk[emptyBlockLoc-i] = harddisk[currentHDIndex+i]
        harddisk[currentHDIndex+i] = EMPTYBLOCK
      }
    }
    for ;currentHDIndex>=0 && harddisk[currentHDIndex] == EMPTYBLOCK; currentHDIndex--{}
  }

  sum := int64(0)
  for i := 0;i<len(harddisk); i++{
    if harddisk[i] > 0 {
      sum += (int64(i) * harddisk[i])
    }
  }
  PrintHDD(harddisk)
  fmt.Println(sum)
}

func PrintHDD(harddisk []int64){
  fmt.Println(harddisk[:100])
  fmt.Println(harddisk[len(harddisk)-50000:len(harddisk)-49900])
}
