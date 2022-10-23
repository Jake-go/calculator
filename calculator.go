// Author: Jacob Golle
// Date: Oct 23, 2022
// Last Revision: Oct 23, 2022
// Description: CLI Calculator for 2 user inputs

package main

import (
	"fmt"
)

func addition(x int, y int) int {
	return x + y
}

func subtraction(x int, y int) int {
	return x - y
}

func multiplication(x int, y int) int {
	return x * y
}

func division(x int, y int) int {
	divResult := x / y
	return int(divResult)
}

func modulus(x int, y int) int {
	return x % y
}

func main() {
	fmt.Println("Hello. This is a calculator app for 2 integer inputs. Please select one of the options to begin:")

	//Variable initialization
	var exit = true
	var userInput int
	var x int
	var y int
	var result int
	for exit {
		//Menu options
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Modulus")
		fmt.Println("6. Exit Program")

		//Obtain menu selection
		//TODO: Sanitize input (create loop -- implement interface?)
		fmt.Scanf("%v", &userInput)
		if userInput == 6 {
			break
		}
		//Obtain user input for 2 numbers
		//TODO: Sanitize input (create loop -- implement interface?)
		fmt.Println("Please input 2 numbers to perform a function on")
		fmt.Scanf("%v %v", &x, &y)

		switch userInput {
		case 1:
			result = addition(x, y)
		case 2:
			result = subtraction(x, y)
		case 3:
			result = multiplication(x, y)
		case 4:
			result = division(x, y)
		case 5:
			result = modulus(x, y)
		case 6:
			exit = false
		}

		fmt.Printf("The result is %v\n", result)
	}
}
