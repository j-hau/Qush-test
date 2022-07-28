package main

import "fmt"

func main() {

	/*
		The goal is if when passed a positive integer, returns true if and only if the decimal representation of the integer contains no odd digits.
		Otherwise return false.
	*/
	positiveInteger := 888888 // Edit the positive integer here
	response := positiveIntegerDecimalChecker(positiveInteger)
	fmt.Println(response)
}

func positiveIntegerDecimalChecker(integer int) bool {
	if integer <= 0 {
		fmt.Println("The input was less than or equal to zero.")
		return false
	}

	for integer > 0 {
		/*
			A pre-requisite to this section is that "integer%10" will return the last character of the integer.
			This last digit can then be removed by dividing the integer by 10.

			This allows us to examine one integer at a time
		*/
		digit := integer % 10
		if (digit % 2) != 0 {
			return false
		}

		integer = integer / 10 // This removes the last number
	}

	return true
}
