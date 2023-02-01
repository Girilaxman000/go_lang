//in go everything is organized in a package

package main

import (
	"fmt"
	"strconv"
)

//basic data types

// <-----  Declaring with Var for dynamic values in a variable ---->

// boolean
// takes false for unassigned value
var myBoolean bool

// string
// takes "" for unassigned value
var myString string

// integer
var myInt, otherInt int

// <-----  Declaring with Const for constant values in a variable ---->

const myConstInt int = 10

func main() {

	//mathematical operations
	myInt = 10
	myInt = myInt * 10
	myInt += 1
	myBoolean = myInt < otherInt

	//string concatenation
	myWorldString := "HelloWorld"
	myHelloWorldString := "Hello" + "" + myWorldString
	fmt.Println(myHelloWorldString)

	//string conversion
	i := 42
	s := strconv.Itoa(i)
	fmt.Println("Integer:", i)
	fmt.Println("String:", s)
	fmt.Println(myConstInt)

	input := "123"
	value, error := strconv.Atoi(input)
	//error handling
	if error != nil {
		return
	}
	fmt.Println(value) // Output: 123
}
