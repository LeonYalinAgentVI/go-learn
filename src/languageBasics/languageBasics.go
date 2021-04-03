package languageBasics

import (
	// "fmt"
	// "log"
	"github.com/LeonYalinAgentVI/go-learn/src/util"
)

func gettingStarted() {
	util.PrintCmd("Getting started", `
	Basic "Hello world" - fmt.Println("Hello, World")
	To run a program, just run 'go run main.go' from the cmd.
	To compile a program, run 'go build -o hw main.go' from the cmd. That will create a 'hw' binary file. To run it, type './hw' from the cmd.
	`)
}

// var packageVar = 2

func variablesAndFunctions() {
	util.PrintCmd("Variable & Functions", `
	- Like in other languages, we have a couple of primitive types:

	var name string = "Leon Yalin"
	var num int = 0

	func sayHello(mas string) {
		fmt.Println(msg)
	}

	var msg string = "Leon Yalin" // use variables
	fmt.Println(sayHello(msg))

	func returnsTwoValues(msg string) (string, string) {
		return msg, "lala"
	}

	// returning multiple values, _ for omitting variables
	var firstVal, _ = returnTwoValues("lulu")
	fmt.Println(firstVal)

	We can have a variable on a package scope, like in Javascript.
	`)
}

func pointers() {
	util.PrintCmd("Pointers", `
	In go, we have pointers and references. References can be used to to get variable memory addresses.
	Pointers are used to get the value stored in particular memory address (reference).

	// Example: change string value using pointers (by reference)
	var text string = "Hello, Bro"
	log.Println("Printing a reference to a variable:", &text)
	var newText *string = &text
	*newText = "Bye, Broh"
	log.Println(*newText, text)
	`)
}

func typesAndStructs() {
	util.PrintCmd("Types and structs", `
	
	`)
}

func LanguageBasics() {
	gettingStarted()
	variablesAndFunctions()
	pointers()
	typesAndStructs()
}
