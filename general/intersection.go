package main

import "fmt"

func main() {
	/*
		The goal here is to find the unique intersection of the below two lists

		when passed two sorted arrays of integers returns an array of any numbers which appear in both.
		No value should appear in the returned array more than once.
	*/

	// Edit the two lists here.
	a := []int{1, 1, 2, 2, 2, 3, 4, 5, 5}
	b := []int{1, 1, 2, 4, 5, 5, 5, 5, 6, 7, 8, 9, 9, 10, 11}

	intersection := intersectionOfLists(a, b)
	fmt.Println("Intersection: ", intersection)
}

func intersectionOfLists(list1 []int, list2 []int) []int {
	// Creating a hash map to get only the unique values from one list
	valueHashMap := make(map[int]bool)

	// Creating a hash map to store if the value is already inside of the intersection
	existsInIntersection := make(map[int]bool)

	var intersection []int // Initialising the empty list for intersection

	for _, i := range list1 {
		valueHashMap[i] = true // Going through the first list and adding each unique element into the map along with 'true' for comparison
	}

	for _, i := range list2 {
		// If the value inside of list 2 exists inside of the unique value hash map of list 1, then append this value to the 'intersection' list which stores the common values.
		if valueHashMap[i] {
			/* 	Now we know the value is common between the lists, we compare it to the hash map of values that already exist in the intersection
			If the value is not already inside of the hash map, append it to the array of intersections and add the value to the hash map so it will not be added to the array of intersections in the future.
			*/
			if !existsInIntersection[i] {
				intersection = append(intersection, i)
				existsInIntersection[i] = true
			}
		}
	}

	return intersection
}
