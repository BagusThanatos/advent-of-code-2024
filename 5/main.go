package main


import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
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
  lines = lines[:len(lines)-1]

  sum := int64(0)

  rulesBefore := make(map[string][]string)

  // parse rules
  i := 0
  for ;lines[i] != "" && i<len(lines);i++ {
    pages := strings.Split(lines[i], "|")
    if v, ok := rulesBefore[pages[0]]; ok {
      rulesBefore[pages[0]] = append(v, pages[1])
    } else {
      rulesBefore[pages[0]] = []string{pages[1]}
    }
  }

  // choose pages
  i+=1
  for ;i<len(lines);i++ {
    valid := true
    pages := strings.Split(lines[i], ",")
    for j:=len(pages)-1;valid && j>=0;j--{
      page := pages[j]
      befores := rulesBefore[page]
      for k:=j-1;valid && k>=0;k--{
        for x:=0;valid && x<len(befores);x++{
          valid = befores[x] != pages[k]
        }
      }
    }

    if valid {
      value, _ := strconv.ParseInt(pages[len(pages)/2], 10, 64)
      sum += value
    }
  }

  fmt.Println(sum)
}