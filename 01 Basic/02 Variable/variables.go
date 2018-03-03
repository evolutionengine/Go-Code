//Go programs start with declaring a package, Package "main" is required for standalone applications
//We can also create custom packages, which we will cover in coming chapters
package main

//Importing the package "fmt" which stands for "format"
//Package "fmt" defines many methods which we can readily use in our program
//Package "fmt" is a part for Standard Library for Go Language
//Other packages in the standard library can be viewed at https://golang.org/pkg/
import "fmt"

//This is a syntax for comment, the compiler ignores comments while compiling

/* This is a syntax for a multi line comment,
Comments are important for explaining your code and
also helpful for other developers working on the project
 */

// Here we declare our "Main" function, this is also the starting point of a program
//The compiler always looks for main.Main() (It means "main" package and "main function") for executing the standalone program

//Go is a functional language, all the pieces of code are wrapped in "functions", it helps to group together all the code
//and provides better readability.

//We can also define our custom functions, which we would be covering in the coming chapters
func main() {
	//	Here we would be learning how to defines variables in Go
	// General syntax -> var variableName dataType
	var name string
	name = "Golang"
	//Print the variable
	fmt.Println(name)
	//Above code simply means, "Print the variable 'name' using the 'Println' function defined in the package 'fmt' "
	//To view more about fmt package log on to https://golang.org/pkg/fmt

	// Similarly for declaring integers
	var number int
	number = 13
	fmt.Println(number)

	// Variables are mutable, i.e their data can be changed
	name = "Go Language"
	fmt.Println(name)
	//This prints out "Go Language" instead of "Golang"

	//We can also define variable for a boolean type
	var whichCondition bool
	whichCondition = true
	fmt.Println(whichCondition)

	//Multiple variables can be defined with the following syntax
	var (
		hobby string
		age int
	)
	hobby = "Programming"
	age = -100
	fmt.Println(hobby)
	fmt.Println(age)

	//Constants are non-mutable variables, i.e data cannot be changed
	const KEY int = 1000
	fmt.Println(KEY)
	//trying to change the value of key with, key = 200, will produce an error

	//Short hand method for declaring variables
	ocean := "Atlantic"
	miles := 1300
	condition := true
	const API_KEY = "7489-2759-2345-9832"
	//Its a common convention to write constants in all capitals for better readability.

	fmt.Println(ocean)
	fmt.Println(miles)
	fmt.Println(condition)
	fmt.Println(API_KEY)
}

/*To run the program follow the following steps
1) Open your terminal and go to the directory which contains the file variables.go
2) On your terminal prompt enter "go run variables.go"
3) For building a standalone executable file enter "go build variables.go"
4) You can run the executable file by typing "./variables"
 */

