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

  enabled := true
  for _, line := range lines {
    boolFirstNumber := false
    boolSecondNumber := false
    firstNum := ""
    secondNum := ""
    for i:=0;i<len(line); {
      char := line[i]
      if char == 'd' && line[i+1] == 'o' && line[i+2] == '(' && line[i+3] == ')'{
        enabled = true
        i+=4
      } else if char == 'd' && line[i+1] == 'o' && line[i+2] == 'n' && line[i+3] == '\'' && line[i+4]=='t' && line[i+5] == '(' && line[i+6]==')'{
        enabled = false
        i+=7
      } else if char == 'm' && line[i+1] == 'u' && line[i+2] == 'l' && line[i+3] == '(' {
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
        if enabled {
          parsedFirst, _ := strconv.ParseInt(firstNum, 10, 64)
          parsedSecond, _ := strconv.ParseInt(secondNum, 10, 64)
          fmt.Println(firstNum, secondNum)
          sum += (parsedFirst * parsedSecond)
        }

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