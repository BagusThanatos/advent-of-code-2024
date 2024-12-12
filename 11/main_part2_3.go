package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)

// NOTE: This also doesn't work, use too much memory and takes too long

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
  split := strings.Split(lines[0], " ")
  line := make([]int64, len(split))
  for i:=0;i<len(split);i++{
    temp, _ := strconv.ParseInt(split[i], 10, 64)
    line[i] = int64(temp)
  }
  count := int64(0)
  for index := 0;index<len(line);index++{

    // fmt.Println(len(line), len(newLine), len(newLine) - len(line), cap(newLine))
    // count += int64(len(newLine))
    count += recursive(line[index], 0)
  }
}

func getMultiplesOfTen(val int64) int64{
  result := int64(0)
  for val > 0 {
    result++
    val /= 10
  }
  if result & 1 != 0 {
    return 0
  }
  temp := int64(1)
  for result>1{
    temp *=10
    result -= 2
  }
  return temp
}

var (
  MAXDEPTH = int(75)
  THRESHOLDDEPTH = int(45)
)
func recursive(val int64, depth int) int64 {
  if MAXDEPTH - depth == THRESHOLDDEPTH {
    newLine := make([]int64, 1)
    newLine [0] = val
    for i:=0;i<THRESHOLDDEPTH;i++{
      x := 0
      border := len(newLine)
      for x<border {
        val := newLine[x]
        if val == 0 {
          newLine[x] = 1
        } else if converted := getMultiplesOfTen(val); converted >0 {
          firstNum := val/converted
          secondNum := val - (firstNum*converted)
          newLine[x] = int64(firstNum)
          newLine = append(newLine, int64(secondNum))
        } else {
          newLine[x] = val * 2024
        }
        x++
      }
    }
    return int64(len(newLine))
  }
  if val == 0 {
    return recursive(1, depth+1)
  }
  if converted := getMultiplesOfTen(val); converted >0 {

    firstNum := val/converted
    secondNum := val - (firstNum*converted)
    return recursive(firstNum, depth+1) + recursive(secondNum, depth+1)
  }
  return recursive(val*2024, depth+1)
}