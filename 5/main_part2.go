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
  checkValid := func (pages []string) bool {
    valid := true
    for j:=len(pages)-1;valid && j>=0;j--{
      page := pages[j]
      befores := rulesBefore[page]
      for k:=j-1;valid && k>=0;k--{
        for x:=0;valid && x<len(befores);x++{
          valid = befores[x] != pages[k]
        }
      }
    }
    return valid
  }
  i+=1
  for ;i<len(lines);i++ { // funny story: previously failed due to putting a -1 here >.>
    pages := strings.Split(lines[i], ",")
    valid := checkValid(pages)

    if !valid {
      // reorder it
      fmt.Println(pages)
      for !checkValid(pages) {
        for j:=len(pages)-1;j>=0;j--{
          page := pages[j]
          befores := rulesBefore[page]
          earliestPos := j
          for k:=j-1;k>=0;k--{
            for x:=0;x<len(befores);x++{
              if befores[x] == pages[k] {
                earliestPos = k
                break
              }
            }
          }
          temp := pages[earliestPos]
          pages[earliestPos] = pages[j]
          pages[j] = temp
        }
      }
      if !checkValid(pages) {
        fmt.Println("still invalid: ", pages)
      }

      value, _ := strconv.ParseInt(pages[len(pages)/2], 10, 64)
      sum += value
    }
  }

  fmt.Println(sum)
}