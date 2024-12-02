package main


import (
  "strconv"
  "os"
  "fmt"
  "io"
  "strings"
  // "sort"
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
  reports := make([][]int64, len(lines-1))
  for i, line := range lines[:len(lines)-1] {
    nums := strings.Split(line, " ")
    levels := make([]int64, len(nums))
    for i:=0;i<len(nums);i++{
      levels[i] = strconv.ParseInt(nums[i], 10, 64)
    }
    reports[i] = levels
  }

  validCount := 0
  for _, line := range reports {
    valid := true
    ascending := true

    num1, _ := line[0]
    num2, _ := line[1]
    if num2 < num1 {
      ascending = false
    }

    if ascending {
      for i:=1; valid && i<len(nums); i++ {
        num1, _ := strconv.ParseInt(nums[i-1], 10, 64)
        num2, _ := strconv.ParseInt(nums[i], 10, 64)
        if num1>=num2 || num2-num1 > 3{
          valid = false
        }
      }
    } else {
      for i:=1; valid && i<len(nums); i++ {
        num1, _ := strconv.ParseInt(nums[i-1], 10, 64)
        num2, _ := strconv.ParseInt(nums[i], 10, 64)
        if num1<=num2 || num1-num2 > 3 {
          valid = false
        }
      }
    }
    if !valid {
      // try by removing one element
    for _, line := range lines[:len(lines)-1] {
      valid := true
      ascending := true
      once := true
      nums := strings.Split(line, " ")

      num1, _ := strconv.ParseInt(nums[0], 10, 64)
      num2, _ := strconv.ParseInt(nums[1], 10, 64)
      if num2 < num1 {
        ascending = false
      }

      if ascending {
        for i:=1; valid && i<len(nums); i++ {
          num1, _ := strconv.ParseInt(nums[i-1], 10, 64)
          num2, _ := strconv.ParseInt(nums[i], 10, 64)
          if num1>=num2 || num2-num1 > 3{
            if i+1 < len(nums) {
              num3, _ := strconv.ParseInt(nums[i+1], 10, 64)
              if !once || (num1>=num3 || num3-num1 > 3){
                valid = false
              } else {
                i++
                once = false
              }
            } else {
              valid = false
            }
          }
        }
      } else {
        for i:=1; valid && i<len(nums); i++ {
          num1, _ := strconv.ParseInt(nums[i-1], 10, 64)
          num2, _ := strconv.ParseInt(nums[i], 10, 64)
          if num1<=num2 || num1-num2 > 3 {
            if i+1 < len(nums) {
              num3, _ := strconv.ParseInt(nums[i+1], 10, 64)
              if !once || (num1>=num3 || num3-num1 > 3) {
                valid = false
              } else {
                i++
                once = false
              }
            } else {
              valid = false
            }
          }
        }
      }
    }
    if valid && num1 != num2 {
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