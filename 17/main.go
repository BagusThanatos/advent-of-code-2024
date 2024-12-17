package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  // "math"
  "strconv"
)

var lines []string
var data [][]int64

var (
  registerA = int64(0)
  registerB = int64(0)
  registerC = int64(0)
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
  lines = strings.Split(string(bytes), "\n")
  lines = lines[:len(lines)-1]

  registerA, _ = strconv.ParseInt(strings.Split(lines[0], ":")[1][1:], 10, 64)
  registerB, _ = strconv.ParseInt(strings.Split(lines[1], ":")[1][1:], 10, 64)
  registerC, _ = strconv.ParseInt(strings.Split(lines[2], ":")[1][1:], 10, 64)
  fmt.Println(registerA, registerB, registerC)

  program := make([]byte, 0)
  for _, v:= range strings.Split(strings.Split(lines[4], ":")[1][1:], ","){
    program = append(program, v[0])
  }
  counter := int(0)
  output := make([]int64, 0)
  for counter < len(program) {
    opcode := program[counter]
    operand := int(program[counter+1] - '0')
    fmt.Println(counter, string(opcode), operand)
    switch opcode{
    case '0':
      registerA = registerA/Pow(2,getComboOperand(operand))
    case '1':
      registerB = (registerB ^ int64(operand))
    case '2':
      registerB = int64(getComboOperand(operand) %8)
    case '3':
      if registerA != 0 {
        fmt.Println("jnz", counter, operand, operand -2 )
        counter = int(operand - 2)
      }
    case '4':
      registerB = registerB ^ registerC
    case '5':
      output = append(output, int64(getComboOperand(operand) % 8))
    case '6':
      registerB = registerA/Pow(2,getComboOperand(operand))
    case '7':
      registerC = registerA/Pow(2, getComboOperand(operand))
    }
    counter += 2
  }
  for _, v := range output{
    fmt.Printf("%d,", v)
  }
  fmt.Println()
}
func getComboOperand(operand int) int64 {
  if operand >= 0 && operand <= 3 {
    return int64(operand)
  }
  if operand == 4 {
    return registerA
  }
  if operand == 5 {
    return registerB
  }
  return registerC
}
func Pow(a, b int64) int64 {
  result := int64(1)
  for i := int64(0);i<b;i++ {
    result = result * a
  }
  return result
}