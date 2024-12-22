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
  // file, err := os.Open("input_coba.txt")
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
  lines = lines[:len(lines)-1]

  changes := make([][]int, len(lines))
  bananas := make([][]int, len(lines))
  for index := range lines {
    secret, _ := strconv.ParseInt(lines[index], 10, 64)
    changes[index] = make([]int, 2000)
    bananas[index] = make([]int, 2000)
    for i:=0;i<2000;i++{
      /*
      Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
  Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
  Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
      */
      prev_secret := secret % 10
      secret = (secret ^ (secret * 64)) % 16777216
      secret = (secret ^ (secret / 32)) % 16777216
      secret = (secret ^ (secret * 2048)) % 16777216
      changes[index][i] = int(secret % 10 - prev_secret)
      bananas[index][i] = int(secret % 10)
    }
  }
  maxBananas := int(0)
  for i:=3;i<len(changes[0]);i++{
    tempBananas := bananas[0][i]
    for m :=1;m<len(changes);m++{
      for n:=3;n<len(changes[m]);n++{
        if sliceEqual(changes[0][i-3:i+1], changes[m][n-3:n+1]){
          tempBananas += bananas[m][n]
          // fmt.Println(tempBananas)
          break
        }
      }
    }
    if tempBananas > maxBananas {
      maxBananas = tempBananas
    }
  }
  fmt.Println(maxBananas)
}
func sliceEqual(a, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  for i:=0;i<len(a);i++ {
    if a[i]!=b[i] {
      return false
    }
  }
  // fmt.Println(a, b)
  return true
}