package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	/*
		When passed a decimal digit X, returns the value of X+XX+XXX+XXXX.
		E.g. if the supplied digit is 3 it should return 3702
		(3+33+333+3333).
	*/

	inputDigit := 6 // Input the digit to test here

	sum, err := increasingQuantitySum(inputDigit) // The function that carries out the addition

	// If there was an error, print the error.  Otherwise print the sum.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}

func increasingQuantitySum(digit int) (int, error) {
	// The initial check to see if it is a decimal representation.  It needs to be 0-9 as a constraint.
	if digit < 0 || digit > 9 {
		return 0, errors.New("The input digit is not a digit.")
	}

	// Initialising the sum to zero.
	sum := 0

	// In the specification of the problem, we want to loop from X -> XXXX and therefore we need to loop 4 times.
	for i := 1; i < 5; i++ {

		// We first get a string representation of the digit that has been validated to be a decimal represented digit.
		digitAsString := strconv.Itoa(digit)

		/*
			Depending on the iteration of the loop, we want to repeat this stringed decimal I times.
			E.g. on the i = 3 loop, we want to make the stringed digit "3" become "333".
		*/
		multipleDigitString := strings.Repeat(digitAsString, i)

		// We then convert this string of multiple digits back to an integer and catch any errors.
		multipleDigit, err := strconv.Atoi(multipleDigitString)

		if err != nil {
			return 0, err
		}

		// Should there be no errors, we simply add the multiple "digit" integer to the sum.
		sum += multipleDigit

	}
	return sum, nil
}
