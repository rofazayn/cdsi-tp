package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i int64
	var j int64
	var sum float64
	var tva float64
	var ttc float64

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How many products are there?")
	scanner.Scan()
	i, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	for j = 0; j < i; j++ {
		fmt.Printf("Give the name of the product number %d: \n", j+1)
		scanner.Scan()
		input := scanner.Text()
		fmt.Println("Type in the cost of the product!")
		scanner.Scan()
		input2, _ := strconv.ParseFloat(scanner.Text(), 10)

		fmt.Printf("Product name: %s , HT price: %f \n", input, input2)
		sum += input2
	}
	tva = (sum * 20) / 100
	ttc = tva + sum
	fmt.Printf("Subtotal: %f \n + TVA %f ,\n Total: %f \n", sum, tva, ttc)
}
