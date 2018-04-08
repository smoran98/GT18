package main

import (
	"fmt"

	util "./package"
)

func main() {

	fmt.Print("1 for Infix\n")
	fmt.Print("2 for postfix\n")
	fmt.Print("3 To Exit\n")

	//option
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		//this here asks the user for infix
		fmt.Println("Option", input, "Was Entered !!!")
		fmt.Print("Enter infix expression: ")
		infix, err := util.ReadUserInput()

		// Error handles data
		if err != nil {
			fmt.Println("Error when scanning input:", err.Error())
			return
		}

		fmt.Println("infix", infix)
		//here
		postfix := util.IntoPost(infix)
		fmt.Println("postfix notation:", postfix)

		fmt.Print("Enter a string for nfa: ")
		stringMatch, err := util.ReadUserInput()

		if err != nil {
			fmt.Println("Error when scanning input:", err.Error())
			return
		}

		fmt.Println(" match ?", util.Pomatch(postfix, stringMatch))

	case 2:
		fmt.Println("Option", input, "Was Entered !!!")
		fmt.Print("Enter postfix expression: ")

		infix, err := util.ReadUserInput()
		// Error handle data
		if err != nil {
			fmt.Println("Error when scanning input:", err.Error())
			return
		}

		fmt.Print("Enter a string for nfa: ")
		stringMatch, err := util.ReadUserInput()

		if err != nil {
			fmt.Println("Error when scanning input:", err.Error())
			return
		}

		fmt.Println("match ?", util.Pomatch(infix, stringMatch))
	default:
		fmt.Println("Enter one of the options")
	}

}
