package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define the pre-defined string to compare to
	preDefinedString := "helloworld"

	fmt.Print("You: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("==> : ")
	input, _ := reader.ReadString('\n')
	var letterStorage []string

	letterStorage = append(letterStorage, input)

	withoutSpace := strings.Replace(input, " ", "", -1)
	// Compare the input to the pre-defined string

	if withoutSpace == "helloworld" {
		fmt.Println("correct")
	}
	if strings.Compare(withoutSpace, preDefinedString) == 0 {
		fmt.Println("The input matches the pre-defined string.")
	} else {
		fmt.Println("The input does not match the pre-defined string.")
	}
	fmt.Println(withoutSpace)
}
