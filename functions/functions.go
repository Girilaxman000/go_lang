package main

import "fmt"

//function declaration
// func greeting(name string) string {
// 	return name
// }

type result struct {
	number int
	text   string
}

func addAndConcat(a int, b int, s string) result {
	return result{number: a + b, text: s}
}

func main() {
	addAndConcat(1, 2, "Hello")
	fmt.Println("laxman")
}
