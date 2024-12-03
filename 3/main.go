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


  sum := int64(0)

  for _, line := range lines {
    boolFirstNumber := false
    boolSecondNumber := false
    firstNum := ""
    secondNum := ""
    for i:=0;i<len(line); {
      char := line[i]
      if char == 'm' && line[i+1] == 'u' && line[i+2] == 'l' && line[i+3] == '(' {
        boolFirstNumber = true
        i += 4
      } else if boolFirstNumber && char <= '9' && char >= '0' {
        boolFirstNumber = true
        firstNum += string(char)
        i++
      } else if boolFirstNumber && char == ',' {
        boolFirstNumber = false
        boolSecondNumber = true
        i++
      } else if boolSecondNumber && char <= '9' && char >= '0' {
        secondNum += string(char)
        i++
      } else if boolSecondNumber && char == ')' {
        parsedFirst, _ := strconv.ParseInt(firstNum, 10, 64)
        parsedSecond, _ := strconv.ParseInt(secondNum, 10, 64)
        fmt.Println(firstNum, secondNum)
        sum += (parsedFirst * parsedSecond)
        i++
        boolFirstNumber = false
        boolSecondNumber = false
        firstNum = ""
        secondNum = ""
      } else {
        // boolInmul = false
        boolFirstNumber = false
        boolSecondNumber = false
        firstNum = ""
        secondNum = ""
        i++
      }
    }
  }
  fmt.Println(sum)
}

func absInt(x int64) int64{
  if (x<0) {
    return -x
  }
  return x
}