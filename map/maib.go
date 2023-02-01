package main

import "fmt"

//struct vs map

func main() {
	//maps are key value pairs just like object in JS.
	//data type for key and value
	emails := make(map[string]string)

	// insert or update a value in a map using the syntax emails[key] = value
	// To retrieve a value, use the syntax value = emails[key]
	//Assign key values
	//len(emails) for map length

	emails["Bob"] = "laxman@gmail.com"
	emails["Bob"] = "laxman@gmail.com"
	fmt.Println(emails)
}
