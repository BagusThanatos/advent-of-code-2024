package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)

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
  for i:=0;i<25;i++{
    newLine := make([]int64, 0)
    for index := 0;index<len(line);index++{
      val := line[index]
      converted := strconv.FormatInt(int64(val), 10)
      if val == 0 {
        newLine = append(newLine, 1)
      } else if len(converted) & 1 == 0 {
        firstNum, _ := strconv.ParseInt(converted[:len(converted)/2], 10, 64)
        secondNum, _ := strconv.ParseInt(converted[len(converted)/2:], 10, 64)
        // fmt.Println(converted, firstNum, secondNum)
        newLine = append(newLine, int64(firstNum), int64(secondNum))
      } else {
        newLine = append(newLine, val*2024)
      }
    }
    fmt.Println(len(line), len(newLine), len(newLine) - len(line), cap(newLine))
    line = newLine
  }


  fmt.Println(len(line))
}