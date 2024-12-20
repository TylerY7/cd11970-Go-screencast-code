package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	make string
	year uint16
	used bool
}

func (c Car) described() string {
	used := ""

	if c.used {
		used = "a used car"
	} else {
		used = "a new car"
	}

	return "This " + strconv.Itoa(int(c.year)) + " " + c.make + " is " + used
}

func main() {

	car1 := Car{make: "DeLorean", year: 1985, used: true}
	car2 := Car{make: "Honda", year: 2023, used: false}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.described())
	fmt.Println(car1.year)
	fmt.Println(car2.described())
	fmt.Println(car2.make)

}
