package main


import (
	"strconv"
	"os"
	"fmt"
	"io"
	"strings"
	"sort"
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

	leftArray := make([]int64, len(lines)-1)
	rightArray := make([]int64, len(lines)-1)
	var sum int64
	sum = 0
	for i, line := range lines[:len(lines)-1] {
		nums := strings.Split(line, " ")
		num1, _ := strconv.ParseInt(nums[0], 10, 64)
		num2, _ := strconv.ParseInt(nums[1], 10, 64)
		leftArray[i] = num1
		rightArray[i] = num2
	}

	sort.Slice(leftArray, func(i, j int) bool{
		return leftArray[i] < leftArray[j]
	})
	sort.Slice(rightArray, func(i, j int) bool{
		return rightArray[i] < rightArray[j]
	})

	for i:=0; i< len(lines)-1; i++{
		sum += absInt(rightArray[i] - leftArray[i])
		fmt.Println(leftArray[i], rightArray[i], sum)
	}
}

func absInt(x int64) int64{
	if (x<0) {
		return -x
	}
	return x
}