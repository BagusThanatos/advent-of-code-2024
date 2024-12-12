package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
  "sync"
)

// try using recursive
var lines []string
var MAXDEPTH = int64(75)
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
  var wg sync.WaitGroup
  result := make([]int64, len(split))

  for i:=0;i<len(split)/2;i++{
    fmt.Println(i)
    wg.Add(1)
    go func (index int) {
      defer wg.Done()
      temp, _ := strconv.ParseInt(split[index], 10, 64)
      result[index] = recursive(int64(temp), 0)
      fmt.Println(result[index])
    }(i)
  }

  wg.Wait()

  for i:=len(split)/2;i<len(split);i++{
    fmt.Println(i)
    wg.Add(1)
    go func (index int) {
      defer wg.Done()
      temp, _ := strconv.ParseInt(split[index], 10, 64)
      result[index] = recursive(int64(temp), 0)
      fmt.Println(result[index])
    }(i)
  }

  wg.Wait()

  count := int64(0)
  for i:=0;i<len(result);i++{
    count += result[i]
  }
  fmt.Println(getMultiplesOfTen(0), getMultiplesOfTen(10), getMultiplesOfTen(100), getMultiplesOfTen(9999), 9999 - (9999/getMultiplesOfTen(9999)))
  fmt.Println(count)
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

func recursive(val, depth int64) int64 {
  if depth >= MAXDEPTH {
    return 1
  }
  if val == 0 {
    return recursive(1, depth+1)
  }
  // converted := strconv.FormatInt(int64(val), 10)
  if converted := getMultiplesOfTen(val); converted >0 {
    // firstNum, _ := strconv.ParseInt(converted[:len(converted)/2], 10, 64)
    // secondNum, _ := strconv.ParseInt(converted[len(converted)/2:], 10, 64)
    // fmt.Println(val)
    firstNum := val/converted
    secondNum := val - (firstNum*converted)
    return recursive(firstNum, depth+1) + recursive(secondNum, depth+1)
  }
  return recursive(val*2024, depth+1)
}