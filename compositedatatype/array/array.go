package main

import (
	"fmt"
	"learn-go/importexport/import_export"
)

//Combine basic types into Composite Types are formed by combining basic types

// array declaration
var exampleArray [3]int

//this type of declaration cann't be used outside of function
//this is an other way of declaring variable
// anotherExampleArray := [3]int{1,2,3}

//slice declaration

// slice and array are same slice vary in length.
var sliceExample []int

//variable already declared := so we can use = operator only. don't get confused

func main() {
	input := "hello"
	output := import_export.Reverse(input)
	fmt.Println(output)

	anotherSliceExample := []string{"ho", "laxm", "giri"}
	anotherExampleArray := [3]int{1, 2, 3}
	anotherExampleArray[0] = 10

	//finding types of variables
	fmt.Printf("%T\n", anotherExampleArray)

	//length of an array
	fmt.Println(len(anotherExampleArray))

	//slice using make()
	sliceArray := make([]int, 10, 15)

	//append to an existing array
	// numbers = append(numbers, i)?
	//cap() for finding capacity

	//increasing capacity by two times

	//1D Array
	numbers := make([]int, 0, 5)

	for i := 0; i < 20; i++ {
		numbers = append(numbers, i)
		fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))
	}

	// numbers := make([]int, 0, 5)

	// for i := 0; i < 20; i++ {
	// 	numbers = append(numbers, i)
	// 	if len(numbers) == cap(numbers) {
	// 		newNumbers := make([]int, len(numbers), 2*cap(numbers))
	// 		copy(newNumbers, numbers)
	// 		numbers = newNumbers
	// 	}
	// 	fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))
	// }

	//2D Array

	fmt.Println(anotherExampleArray[0], anotherSliceExample[0], sliceArray)
}
