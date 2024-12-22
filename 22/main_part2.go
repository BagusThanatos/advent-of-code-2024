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
  // file, err := os.Open("input.txt")
  file, err := os.Open("input_coba.txt")
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

  sum := int64(0)
  for index := range lines {
    secret, _ := strconv.ParseInt(lines[index], 10, 64)
    for i:=0;i<2000;i++{

      /*
      Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
  Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
  Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
      */
      secret = (secret ^ (secret * 64)) % 16777216
      secret = (secret ^ (secret / 32)) % 16777216
      secret = (secret ^ (secret * 2048)) % 16777216
    }
    sum += secret
  }

  fmt.Println(sum)
}
