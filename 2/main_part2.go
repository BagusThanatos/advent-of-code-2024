package main


import (
  "strconv"
  "os"
  "fmt"
  "io"
  "strings"
)


func checkValid(line []int64) bool {
  valid := true
  ascending := true

  num1 := line[0]
  num2 := line[1]
  if num2 < num1 {
    ascending = false
  }

  if ascending {
    for i:=1; valid && i<len(line); i++ {
      num1 := line[i-1]
      num2 := line[i]
      if num1>=num2 || num2-num1 > 3{
        valid = false
      }
    }
  } else {
    for i:=1; valid && i<len(line); i++ {
      num1 := line[i-1]
      num2 := line[i]
      if num1<=num2 || num1-num2 > 3 {
        valid = false
      }
    }
  }
  return valid && num1 != num2
}


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
  reports := make([][]int64, len(lines)-1)
  for i, line := range lines[:len(lines)-1] {
    nums := strings.Split(line, " ")
    levels := make([]int64, len(nums))
    for i:=0;i<len(nums);i++{
      levels[i], _ = strconv.ParseInt(nums[i], 10, 64)
    }
    reports[i] = levels
  }

  validCount := 0
  for _, line := range reports {
    valid := checkValid(line)
    if !valid {
      // try remove 1 element
      for i:=0; !valid && i<len(line); i++ {
        var temp []int64
        temp = append(temp, line[:i]...)
        temp = append(temp, line[i+1:]...)
        valid = checkValid(temp)
      }
    }
    if valid{
      fmt.Println(line)
      validCount++
    }
  }

  fmt.Println(validCount)
}

func absInt(x int64) int64{
  if (x<0) {
    return -x
  }
  return x
}