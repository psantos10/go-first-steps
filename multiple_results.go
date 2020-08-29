package main

import (
	"fmt"
	"strings"
)

func splitFullName(fullname string) (string, string) {
	split := strings.Split(fullname, " ")

	return split[0], split[len(split)-1]
}

func multipleResults() {
	fullName := "Patrício dos Santos"
	firstName, lastName := splitFullName(fullName)

	fmt.Println("Full Name:", fullName)
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
}
