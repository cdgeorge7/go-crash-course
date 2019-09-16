package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key-values
	// emails["bob"] = "bob@gmail.com"
	// emails["biscuits"] = "biscuits@gmail.com"
	// emails["chris"] = "chris@gmail.com"

	// Declare map and add key-values
	emails := map[string]string{"bob": "bob@gmail.com", "biscuits": "biscuits@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
