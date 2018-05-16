package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const UsernameRegex string = `^@?(\w){1,15}$`

func CheckUserNameSyntax(username string) bool {
	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatalln(err)
	}
	validationResult = r.MatchString(username)
	return validationResult
}

func main() {
	var userInput string
	flag.StringVar(&userInput, "username", "Gopher", "Validates username")
	flag.Parse()

	fmt.Println("Username validation checker running...")
	fmt.Println("Checking syntax for", userInput, ", resulted in:", CheckUserNameSyntax(userInput))
}
