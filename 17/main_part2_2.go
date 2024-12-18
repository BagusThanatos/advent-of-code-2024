package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  "strconv"
)

// NOTE: this is THE SOLUTION
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

  // registerA, _ = strconv.ParseInt(strings.Split(lines[0], ":")[1][1:], 10, 64)
  registerB, _ = strconv.ParseInt(strings.Split(lines[1], ":")[1][1:], 10, 64)
  registerC, _ = strconv.ParseInt(strings.Split(lines[2], ":")[1][1:], 10, 64)
  fmt.Println(registerA, registerB, registerC)

  program := make([]byte, 0)
  for _, v:= range strings.Split(strings.Split(lines[4], ":")[1][1:], ","){
    program = append(program, v[0])
  }
  fmt.Println(program)

  // The idea here is to build the output from LAST
  /*
  This is the program
  registerB = registerA%8 < registerB = 1-7
  registerB = registerB ^ int64(1) < 1-7 ^ 1 -> only 7 possibilities
  registerC = registerA/Pow(2, registerB) (1-7)/pow(2, 1-7)
  registerB = registerB XOR registerC < 4= something XOR
  registerA = registerA/Pow(2, 3) < 1-7
  registerB = registerB ^ int64(4) < 0, registerB = 4
  OUTPUT getComboOperand(registerB) <0
  jnz registerA < 0
  */
  // Note that register A is always divided by * at the end, hence we can simplify how we look for the starting point below
  startA := int64(0)
  output := make([]byte, 0)
  MAXOUTPUT := len(program)-1
  for MAXOUTPUT >= 0 && !Equal(program[MAXOUTPUT:], output) {
    startA++
    output = make([]byte, 0)
    registerA = startA
    counter := int(0)
    for counter < len(program) {
      opcode := program[counter]
      operand := int(program[counter+1] - '0')
      switch opcode{
      case '0':
        registerA = registerA/Pow(2,getComboOperand(operand))
      case '1':
        registerB = (registerB ^ int64(operand))
      case '2':
        registerB = int64(getComboOperand(operand) %8)
      case '3':
        if registerA != 0 {
          counter = int(operand - 2)
        }
      case '4':
        registerB = registerB ^ registerC
      case '5':
        output = append(output, byte(getComboOperand(operand) % 8)+'0') // as we're dealing with '0' here in the operand
      case '6':
        registerB = registerA/Pow(2,getComboOperand(operand))
      case '7':
        registerC = registerA/Pow(2, getComboOperand(operand))
      }
      counter += 2
    }
    if Equal(program[MAXOUTPUT:], output){
      fmt.Println(program[MAXOUTPUT:], output, Equal(program[MAXOUTPUT:], output), startA)
      MAXOUTPUT--
      if MAXOUTPUT < 0 {
        break
      }
      startA = (startA * 8) -2
    }
  }
  fmt.Println(program, output, Equal(program, output), startA)
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
func Equal(a, b []byte) bool {
  // fmt.Println(a,b)
  if len(a) != len(b) {
    return false
  }
  for i:=0;i<len(a);i++{
    if (a[i]!=b[i]){
      return false
    }
  }
  return true
}