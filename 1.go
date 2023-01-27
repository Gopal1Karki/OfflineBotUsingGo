package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a map of pre-defined responses
	responses := map[string]string{
		"hi":        "Hello! How can i assist you?",
		"howareyou": "I'm good, thank you!",
		"bye":       "Goodbye!",
	}

	// Continuously ask for input
	for {
		fmt.Print("You: ")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		input = strings.Replace(input, " ", "", -1)

		// Check if the input matches any of the pre-defined responses
		response, ok := responses[input]
		if input == "open" {
			fmt.Println("Opening the app..")
			Open()

		}
		if ok {
			fmt.Println("Bot: ", response)
		} else {
			fmt.Println("Bot: I'm sorry, I don't understand.")
		}

	}

}
func Open() {
	fmt.Println("hello")
	main()
}
