package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// emails["Bob"] = "bob@gmail.com"
	// emails["Fa"] = "fa@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "Fa": "fa@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Fa"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
