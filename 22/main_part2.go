package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)


// Note: This could solve the problem but is not fast enough I think
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
  maxBananas := int64(0)
  mark := make([][][][]bool, 20)
  for i:=0;i<len(mark);i++{
    mark[i] = make([][][]bool, 20)
    for j:=0;j<len(mark[i]);j++{
      mark[i][j] = make([][]bool, 20)
      for k:=0;k<len(mark[i][j]);k++{
        mark[i][j][k] = make([]bool, 20)
      }
    }
  }
  for baseRow := 0;baseRow<len(changes)-1;baseRow++{
    for i:=3;i<len(changes[baseRow]);i++{
      c := changes[baseRow][i-3:i+1]
      if mark[c[0]][c[1]][c[2]][c[3]] {
        continue
      }
      mark[c[0]][c[1]][c[2]][c[3]]= true
      tempBananas := bananas[baseRow][i]
      // tempBananas := int64(0)

      // there's no need to check previous rows as it should've been found previously
      for m := baseRow+1;m<len(changes);m++{
        // if m == baseRow {
        //   continue
        // }
        for n:=3;n<len(changes[m]);n++{
          if sliceEqual(c, changes[m][n-3:n+1]){
            tempBananas += bananas[m][n]
            break
          }
        }
      }
      if tempBananas > maxBananas {
        fmt.Println(baseRow, i, changes[baseRow][i-3:i+1], bananas[baseRow][i], tempBananas)
        maxBananas = tempBananas
      }
    }
  }
  fmt.Println(maxBananas)
}
func sliceEqual(a, b []int64) bool {
  return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}