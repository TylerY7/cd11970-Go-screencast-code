package main

import "fmt"

func main() {
	//var product = "T-shirt"
	//var cost int = 20

	product := "T-shirt"
	cost := 20

	fmt.Println("product's value is:", product)
	fmt.Printf("product type is: %T\n", product)

	cost = 15

	fmt.Println("product's cost is:", cost)
	fmt.Printf("product type is: %T\n", cost)
}
