package main

import "fmt"

func main() {
	/*
		0 1 1 2 3 5 8 13
		var fibonacciSlice []int


		fibonacciArray[0] = 0
		fibonacciArray[1] = 1
		fibonacciArray[2] = 1
		fibonacciArray[3] = 2
		fibonacciArray[4] = 3
		fibonacciArray[5] = 5
		fibonacciArray[6] = 8
		fibonacciArray[7] = 13

		fmt.Println(fibonacciArray)
	*/

	fibonacciSlice := []int{0, 1, 1, 2, 3, 5, 8, 13}
	fmt.Println(fibonacciSlice)

	fmt.Println(fibonacciSlice[0])
	fmt.Println(fibonacciSlice[7])
	fmt.Println(len(fibonacciSlice))
	fmt.Println(fibonacciSlice[0:4])

	fibonacciSlice = append(fibonacciSlice, 21)
	fmt.Println(fibonacciSlice)
	fmt.Println(len(fibonacciSlice))
}
