package main

import (
  "os"
  "fmt"
  "io"
  "strings"
  // "reflect"
  "strconv"
)


// NOTE: This is NOT the solution
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
  // startA := int64(1376984737)
  startA := int64(Pow(8, 15)- 100)
  // startA := int64(1)
  // pow := 1
  output := make([]byte, 0)
  for !Equal(program, output) && startA < Pow(8, 17){
    output = make([]byte, 0)
    // startA = startA * 8
    startA++
    // fmt.Println("startA", startA)
    registerA = startA
    counter := int(0)
    // fmt.Println(">>>> HERE")
    for counter < len(program) {
      opcode := program[counter]
      operand := int(program[counter+1] - '0')
      // fmt.Println(counter, string(opcode), operand)
      switch opcode{
      case '0':
        registerA = registerA/Pow(2,getComboOperand(operand))
        // fmt.Printf("registerA = registerA/Pow(2, combo(%d)\n", operand)
      case '1':
        registerB = (registerB ^ int64(operand))
        // fmt.Printf("registerB = registerB ^ int64(%d)\n", operand)
      case '2':
        registerB = int64(getComboOperand(operand) %8)
        // fmt.Printf("registerB = getComponen(%d) %%8\n", operand)
      case '3':
        if registerA != 0 {
          // fmt.Println("jnz", counter, operand, operand -2 )
          counter = int(operand - 2)
        }
        // fmt.Printf("jnz registerA\n")
      case '4':
        registerB = registerB ^ registerC
        // fmt.Printf("registerB = registerB XOR registerC\n")
      case '5':
        output = append(output, byte(getComboOperand(operand) % 8)+'0') // as we're dealing with '0' here in the operand
        // fmt.Printf("OUTPUT getComboOperand(%d)\n", operand)
      case '6':
        registerB = registerA/Pow(2,getComboOperand(operand))
        // fmt.Printf("registerB = registerA/Pow(2, combo(%d)\n", operand)
      case '7':
        registerC = registerA/Pow(2, getComboOperand(operand))
        // fmt.Printf("registerC = registerA/Pow(2, combo(%d)\n", operand)
      }
      counter += 2
    }
    // fmt.Println(output)
    // break
  }
  // for _, v := range output{
  //   fmt.Printf("%d,", v)
  // }
  fmt.Println(Equal(program, output), startA)
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