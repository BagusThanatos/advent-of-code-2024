package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)


// Note: This is a MUCH FASTER solution
var lines []string
func main() {
  // file, err := os.Open("input.txt")
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
  // fmt.Println(lines[len(lines)-1])
  lines = lines[:len(lines)-1]

  changes := make([][]int64, len(lines))
  bananas := make([][]int64, len(lines))
  for index := range lines {
    secret, _ := strconv.ParseInt(lines[index], 10, 64)
    changes[index] = make([]int64, 2000)
    bananas[index] = make([]int64, 2000)
    for i:=0;i<2000;i++{
      /*
      Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
  Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
  Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
      */
      prevBananas := secret % 10
      secret = (secret ^ (secret * 64)) % 16777216
      secret = (secret ^ (secret / 32)) % 16777216
      secret = (secret ^ (secret * 2048)) % 16777216
      changes[index][i] = (secret % 10) - prevBananas + 9 // convert range from [-9, 9] to [0, 18]
      bananas[index][i] = secret % 10
    }
  }
  mark := make(map[changeNode]int64)
  for baseRow := 0;baseRow<len(changes);baseRow++{
    tempMark := make(map[changeNode]int64)
    for i:=3;i<len(changes[baseRow]);i++{
      c := changes[baseRow][i-3:i+1]
      k := changeNode{c[0],c[1],c[2],c[3]}
      if _, ok := tempMark[k]; !ok {
        tempMark[k] = mark[k] + bananas[baseRow][i]
      }
    }
    // fill back new values
    for k := range tempMark {
      mark[k] = tempMark[k]
    }
  }
  maxBananas := int64(0)
  for k := range mark {
    if mark[k] > maxBananas {
      fmt.Println(k, mark[k])
      maxBananas = mark[k]
    }
  }
  fmt.Println(maxBananas)
}

type changeNode struct{
  a,b,c,d int64
}
func sliceEqual(a, b []int64) bool {
  return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}