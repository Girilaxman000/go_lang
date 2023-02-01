package main

import "fmt"

func main() {
	// range keyword is used to iterate over elements in a variety of data structures, including arrays, slices, strings, maps, and channels.

	ids := []int{1, 2, 3, 4, 5}

	for i, id := range ids {
		fmt.Println(i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Println(id)
	}

}
