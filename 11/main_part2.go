package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)

// NOTE: This is the way to go. Use map to reduce the amount of compute AND memory used
//       The problem has this note `it always retains order`, which in this case actually means `the order does not matter at all as long as you can count`

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
  m := make(map[int64]int64)
  for index := 0;index<len(line);index++{
    m[line[index]] = 1
  }

  fmt.Println(m)
  for i:=0;i<25;i++ {
    newMap := make(map[int64]int64)
    for val := range m {
      if val == 0 {
        newMap[1] =  newMap[1] + m[val]
      } else if converted := getMultiplesOfTen(val); converted >0 {
        firstNum := val/converted
        secondNum := val - (firstNum*converted)
        newMap[firstNum] = newMap[firstNum] + m[val]
        newMap[secondNum] = newMap[secondNum] + m[val]
      } else {
        newMap[val * 2024] += m[val]
      }
    }
    m = newMap
  }
  count := int64(0)
  for _, v := range m{
    count+= v
  }
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