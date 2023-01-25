package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func startInterface() {
	fmt.Printf("\n\n")
	fmt.Println(" GO  OFFLINE  BOT")

	fmt.Printf("\n\n")
	fmt.Println("GoBOT: Hello ! I am GoBOT. How can i assist you?")
	fmt.Printf("\n")
	mainBot()
}

func mainBot() {
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Printf("==> : ")
	sentence, _ := reader.ReadString('\n')
	reply(sentence)*/
	var inp string
	fmt.Printf("==> ")
	fmt.Scanln(&inp)
	reply(inp)

}

func reply(sentence string) {
	inf := sentence
	if inf == "hello" || inf == "hi" || inf == "Hello" || inf == "Hi" || inf == "HELLO" || inf == "HI" {
		fmt.Printf("GoBOT> ")
		fmt.Println(" Hello !")
		mainBot()
	} else if inf == "Time " || inf == "time" || inf == "TIME" {
		now := time.Now()
		fmt.Printf("GoBOT> ")
		if now.Hour() > 12 {
			hr := now.Hour() - 12
			ampm := "PM"
			fmt.Println("Current Time is : ", hr, ":", now.Minute(), ampm)
		} else {
			ampm := "AM"
			fmt.Println("Current Time is : ", now.Hour(), ":", now.Minute(), ampm)
		}

		mainBot()
	} else if inf == "openbrave" {
		fmt.Println("Opening the Brave......")
		time.Sleep(1 * time.Second)
		openBrowser("Brave")
	} else if inf == "date" || inf == "Date" || inf == "Date" {
		now := time.Now()
		fmt.Printf("GoBOT> ")
		fmt.Println("Current Date is : ", now.Day(), "/", now.Month(), "/", now.Year())

	} else if inf == "/cls" || inf == "/clear" {
		fmt.Printf("==> Clearing the screen !!!")
		time.Sleep(2 * time.Second)
		main()
	} else if inf == "exit" || inf == "Exit" || inf == "EXIT" {
		fmt.Println("Exiting the GoBOT.....!")
		time.Sleep(1 * time.Second)
		clearScreen()
		os.Exit(1)

	}
	mainBot()

}
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
func main() {
	clearScreen()
	startInterface()
}
