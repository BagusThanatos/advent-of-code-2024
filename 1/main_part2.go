package main


import (
  "strconv"
  "os"
  "fmt"
  "io"
  "strings"
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

  leftMap := make(map[int64]int64)
  rightArray := make([]int64, len(lines)-1)
  var sum int64
  sum = 0
  for i, line := range lines[:len(lines)-1] {
    nums := strings.Split(line, " ")
    // fmt.Println(nums)

    num1, _ := strconv.ParseInt(nums[0], 10, 64)
    num2, _ := strconv.ParseInt(nums[1], 10, 64)
    leftMap[num1] = 0
    rightArray[i] = num2
    // sum +=  num1 - num2
    // fmt.Println(sum)
  }
  for i:=0; i< len(lines)-1; i++{
    _, ok := leftMap[rightArray[i]]
    if ok {
      leftMap[rightArray[i]] += 1
    }
  }
  for k, v := range leftMap{
    sum += k * v
    fmt.Println(k, v, sum)
  }
}

func absInt(x int64) int64{
  if (x<0) {
    return -x
  }
  return x
}