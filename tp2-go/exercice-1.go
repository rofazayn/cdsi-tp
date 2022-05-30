package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What's the size of your array of numbers? \n")
	scanner.Scan()
	size, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	var arr = make([]int, size)
	var i int64
	for i = 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i+1)
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		arr[i] = int(input)
	}
	fmt.Println("Your array is: ", arr)

	var flag int = 0
	var t int64
	fmt.Printf("The prime numbers in your array are: \n")
	for t = 0; t < size; t++ {
		flag = 0
		for j := 2; j < arr[t]/2; j++ {
			if arr[t]%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 && arr[t] > 1 {
			fmt.Printf("%d ", arr[t])
		}
	}
}
