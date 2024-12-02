package main


import (
	"strconv"
	"os"
	"fmt"
	"io"
	"strings"
	// "sort"
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

	validCount := 0
	for _, line := range lines[:len(lines)-1] {
		valid := true
		ascending := true
		nums := strings.Split(line, " ")

		num1, _ := strconv.ParseInt(nums[0], 10, 64)
    num2, _ := strconv.ParseInt(nums[1], 10, 64)
    if num2 < num1 {
    	ascending = false
    }

    if ascending {
	    for i:=1; valid && i<len(nums); i++ {
	    	num1, _ := strconv.ParseInt(nums[i-1], 10, 64)
    		num2, _ := strconv.ParseInt(nums[i], 10, 64)
    		if num1>=num2 || num2-num1 > 3{
    			valid = false
    		}
	    }
  	} else {
  		for i:=1; valid && i<len(nums); i++ {
	    	num1, _ := strconv.ParseInt(nums[i-1], 10, 64)
    		num2, _ := strconv.ParseInt(nums[i], 10, 64)
    		if num1<=num2 || num1-num2 > 3 {
    			valid = false
    		}
	    }
  	}
		if valid && num1 != num2 {
			fmt.Println(line)
			validCount++
		}
	}

	fmt.Println(validCount)
}

func absInt(x int64) int64{
	if (x<0) {
		return -x
	}
	return x
}