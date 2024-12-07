package main

/*
NOTE: This is the correct answer
The problem was to put an obstacle BEFORE the guard start patrolling
*/

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)

var lines [][]int64
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

  linesString := strings.Split(string(bytes), "\n")
  lines = make([][]int64, len(linesString)-1)
  for i:=0;i<len(lines);i++{
    temp := strings.Split(linesString[i], ":")
    tempTotal, _ := strconv.ParseInt(temp[0], 10, 64)
    numsString := strings.Split(temp[1], " ")[1:] // frist element is always empty string
    tempNums := []int64{tempTotal}
    for _, v := range numsString{
      parsed, _ := strconv.ParseInt(v, 10, 64)
      tempNums = append(tempNums, parsed)
    }
    if len(tempNums) <=2 {
      fmt.Println(i, tempNums)
    }
    lines[i] = tempNums
  }
  // fmt.Println(lines)

  sum := int64(0)
  for i:=0;i<len(lines);i++{
    if tryCalculate(lines[i][1], lines[i][0], i, 2){
      sum+=lines[i][0]
    }
  }
  fmt.Println(sum)
}

func tryCalculate (currentTotal , total int64, currentLine, index int) bool {
  if index >= len(lines[currentLine]) {
    return currentTotal == total
  }
  return tryCalculate(currentTotal * lines[currentLine][index], total, currentLine, index+1) || tryCalculate(currentTotal + lines[currentLine][index], total, currentLine, index+1)
}
