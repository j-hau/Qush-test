package main

import "fmt"

func main() {
	/*
		 	The goal is to calculate the sum of the first 100 even-valued fibonacci numbers where fibonacci number is denoted by f(n) and:
			f(0) = 1
			f(1) = 1
			f(n) = f(n-1) + f(n-2)
	*/
	sum := evenFibonacciSum()
	fmt.Println("The sum of the first 100 even Fibonacci Numbers is ", sum, "!")
}

func evenFibonacciSum() uint64 {
	/*
		'countOfEven' is keeping count of the number of even fibonacci numbers added to the total sum
		'a' is keeping track of the f(n-2) aspect of f(n)
		'b' is keeping track of the f(n-1) aspect of f(n)
		'sum' is the running total of the even fibonacci numbers
	*/
	var countOfEven int = 0
	var a uint64 = 0
	var b uint64 = 1
	var sum uint64 = 0

	// Iterating for the first 100 even numbers
	for countOfEven < 100 {
		/*
			fibNum is the sum of 'a' and 'b' which represents f(n) = f(n-1) + f(n-2).
			While f(0) = 1, it does not matter in this case as we are not tracking how many fibonacci numbers have been calculated, but tracking how many were even.
		*/
		fibNum := a + b
		/*
			If the calculated fibonacci number is even - worked out via modular division then add the even fibonacci number to the running sum.
			Then increment the counter for the number of even fibonacci numbers
		*/
		if (fibNum % 2) == 0 {
			sum += fibNum
			countOfEven++
		}
		/*
			Then shift the numbers downwards so f(n-2) is now the current f(n-1)
			f(n-1) is the current f(n).

			Repeat this until the counter for even fibonacci numbers reaches 100.
		*/
		a = b
		b = fibNum
	}
	return sum
}
