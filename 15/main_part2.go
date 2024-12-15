package main

import (
  "os"
  "fmt"
  "io"
  "strings"
)

var lines []string
var mapBox [][]byte
func main() {
  file, err := os.Open("input_coba3.txt")
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
  border := 0
  for lines[border] != ""{border++}
  mapBox = make([][]byte, border)
  // make everything twice as wide, except the robot
  for i:=0;i<border;i++{
    mapBox[i] = make([]byte, 0)
    for j:=0;j<len(lines[i]);j++{
      if lines[i][j] == '@' {
        mapBox[i] = append(mapBox[i], '@', '.')
      } else if lines[i][j] == 'O' {
        mapBox[i] = append(mapBox[i], '[', ']')
      }else {
        mapBox[i] = append(mapBox[i], lines[i][j], lines[i][j])
      }
    }

  }
  for _, v := range lines[:border]{
        fmt.Println(string(v))
      }
  for _, v := range mapBox{
        fmt.Println(string(v))
      }
  // look for robot position
  y := 1
  x :=1
  for ;y<len(mapBox)-1;y++{
    x = 1
    for ;x<len(mapBox[y])-1;x++{
      if mapBox[y][x] == '@' {
        goto found
      }
    }
  }; return
  found:
  for _, line := range lines[border+1:] {
    for _, move := range line {
      if move == '<' {
        // look for '.'
        i:=x
        for ;i>=0&& mapBox[y][i] != '#' && mapBox[y][i] != '.';i--{}
        // we can move stuffs around
        if mapBox[y][i] == '.' {
          for ;i<x;i++ {
            mapBox[y][i] = mapBox[y][i+1]
          }
          // robot has moved
          mapBox[y][x] = '.'
          x--
        }
      } else if move == '>' {
        // look for '.'
        i:=x
        for ;i<len(mapBox[y]) && mapBox[y][i] != '#' && mapBox[y][i] != '.';i++{}
        // we can move stuffs around
        if mapBox[y][i] == '.' {
          for ;i>x;i-- {
            mapBox[y][i] = mapBox[y][i-1]
          }
          // robot has moved
          mapBox[y][x] = '.'
          x++
        }
      } else if move == 'v' {
        // look for '.'
          // fmt.Println("here", string(mapBox[y][x]))

        i:=y+1
        if mapBox[i][x] == '.' {
          mapBox[i][x] = '@'
          mapBox[y][x] = '.'
          y++
        } else if mapBox[i][x] == '[' {
          for ;i<len(mapBox) && mapBox[i][x] != '#' && mapBox[i][x] != '.' && mapBox[i][x] == '[' && mapBox[i][x+1] == ']';i++{}
          // we can move stuffs around
          if mapBox[i][x] == '.' && mapBox[i][x+1] == '.' {
            for ;i>y+1;i-- {
              mapBox[i][x] = mapBox[i-1][x]
              mapBox[i][x+1] = mapBox[i-1][x+1]
            }
            mapBox[i][x] = '@'
            // robot has moved
            mapBox[i][x+1] = '.'
            i++
            mapBox[y][x] = '.'
            y++
          }
        } else if mapBox[i][x] == ']' {
          for ;i<len(mapBox) && mapBox[i][x] != '#' && mapBox[i][x] != '.' && mapBox[i][x-1] == '[' && mapBox[i][x] == ']';i++{}
          // we can move stuffs around
          if mapBox[i][x] == '.' && mapBox[i][x-1] == '.' {
            for ;i>y+1;i-- {
              mapBox[i][x] = mapBox[i-1][x]
              mapBox[i][x-1] = mapBox[i-1][x-1]
            }
            mapBox[i][x] = '@'
            // robot has moved
            mapBox[i][x-1] = '.'
            i++
            mapBox[y][x] = '.'
            y++
          }
        }
      } else if move == '^' {
        // look for '.'
        i:=y-1
        if mapBox[i][x] == '.' {
          mapBox[i][x] = '@'
          mapBox[y][x] = '.'
          y--
        } else if mapBox[i][x] == '[' {
          for ;i>=0 && mapBox[i][x] != '#' && mapBox[i][x] != '.' && mapBox[i][x] == '[' && mapBox[i][x+1] == ']';i--{}
          // we can move stuffs around
          if mapBox[i][x] == '.' && mapBox[i][x+1] == '.' {
            for ;i<y-1;i++ {
              mapBox[i][x] = mapBox[i+1][x]
              mapBox[i][x+1] = mapBox[i+1][x+1]
            }
            mapBox[i][x] = '@'
            // robot has moved
            mapBox[i][x+1] = '.'
            i--
            mapBox[y][x] = '.'
            y--
          }
        } else if mapBox[i][x] == ']' {
          for ;i>=0 && mapBox[i][x] != '#' && mapBox[i][x] != '.' && mapBox[i][x-1] == '[' && mapBox[i][x] == ']';i--{}
          // we can move stuffs around
          if mapBox[i][x] == '.' && mapBox[i][x-1] == '.' {
            for ;i<y-1;i++ {
              mapBox[i][x] = mapBox[i+1][x]
              mapBox[i][x-1] = mapBox[i+1][x-1]
            }
            mapBox[i][x] = '@'
            // robot has moved
            mapBox[i][x-1] = '.'
            i--
            mapBox[y][x] = '.'
            y--
          }
        }
      }
      fmt.Println(string(move))
      for _, v := range mapBox{
        fmt.Println(string(v))
      }
    }
  }

  fmt.Println(sumGPSCoordinatesLargerBoxes())

}
func sumGPSCoordinatesLargerBoxes() int64{
  sum := int64(0)
  for i:=1; i<len(mapBox);i++{
    for j:=1;j<len(mapBox[i]);j++{
      if mapBox[i][j]=='['{
        sum += 100*int64(i) + int64(j)
        j++
      }
    }
  }
  return sum
}

