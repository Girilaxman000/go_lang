package main

import "fmt"

func main() {
	//ponters allow to point to the memory address/location of the value that's in a variable
	//where value is stored in memory
	a := 5
	//assigning b to a pointer of a
	b := &a

	//read value from address
	// The * operator is used to dereference the pointer and access the value stored at its memory address.
	fmt.Println(*b, *&a)

	*b = 10
	fmt.Println(a)
}
