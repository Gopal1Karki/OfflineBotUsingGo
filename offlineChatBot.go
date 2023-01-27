package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	var input string
	fmt.Printf("==> ")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "open" {
		OpenApp()
	}
	reply(input)

}

func reply(sentence string) {
	inf := sentence
	responses := map[string]string{
		"hi":    "Hello! How can i assist you?",
		"hello": "Hello! How can i assist you?",
		"bye":   "Goodbye",
	}
	response, ok := responses[inf]
	if ok {
		fmt.Println("GoBOT: ", response)
		fmt.Printf("\n")
		mainBot()
	} else {
		if inf == "time" {
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
		} else if inf == "mediaplayer" {
			fmt.Println("Opening Media PLayer")
			time.Sleep(1 * time.Second)
			mediaPlayer()
		} else {
			fmt.Println("GoBOT> I'm sorry, I don't understand.")
		}
		fmt.Printf("\n")
		mainBot()
	}

}
func OpenApp() {
	fmt.Println("GoBOT> Following are the appname that you can open are:")
	fmt.Println("Brave ,Calculator,notepad")
	fmt.Println("GoBOT> Enter the AppName you want to open.")
	var input string
	fmt.Printf("==> ")
	fmt.Scanln(&input)
	input = strings.ToLower(input)

	if input == "brave" {
		fmt.Println("Opening the Brave......")
		time.Sleep(1 * time.Second)
		openBrowser("Brave")
	} else if input == "calculator" {
		fmt.Println("Opening the Calculator")
		time.Sleep(1 * time.Second)
		calculator()
	} else if input == "notepad" {
		fmt.Println("Wait....")
		time.Sleep(1 * time.Second)
		notePad()
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
	fmt.Println("Brave opened Sucessfully!!!")
	return cmd.Start() == nil
}

func calculator() {
	cmd := exec.Command("calc.exe")

	err := cmd.Run()

	if err != nil {
		fmt.Println("Failed to load calculator")
	} else {
		fmt.Println("Calculator opened sucessfully!")
	}
}
func mediaPlayer() {
	cmd := exec.Command("cmd", "/C", "start", "wmplayer", "wmplayer.exe", "microphone:")
	cmd.Start()
}

func notePad() {
	cmd := exec.Command("notepad.exe")
	err := cmd.Start()

	if err != nil {
		fmt.Println("Failed to open Notepad..")
	} else {
		fmt.Println("Notepad is opened Sucessfully!")
	}

}
func main() {
	clearScreen()
	startInterface()
}
