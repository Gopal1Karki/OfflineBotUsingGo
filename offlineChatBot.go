package main

import (
	"bufio"
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
	fmt.Println("Type /help to if you need any help!! ")
	fmt.Printf("\n")
	fmt.Println("GoBOT: Hello ! I am GoBOT. How can i assist you?")

	mainBot()
}

func mainBot() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("==> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	if input == "/help" {
		goHelp()
	}
	reply(input)

}
func goHelp() {
	fmt.Println("GoBOT> Here are the command you can use in GoBOT!")

	cmdList := []string{
		"/cls", "/clear", "/help", "exit", "/shutdown",
	}
	for _, cmdL := range cmdList {
		fmt.Printf("->")
		fmt.Println(cmdL)
	}
	fmt.Printf("\n")
	mainBot()
}
func reply(sentence string) {
	inf := sentence
	responses := map[string]string{
		"hi":                "Hi! How can i assist you?",
		"hello":             "Hello !",
		"what is your name": "My name is GoBOT!",
		"how are you":       "I'm good, thank you!",
		"good morning":      "Goodmorning ! Have a good Day!",
		"goodmorning":       "Goodmorning ! Have a good Day!",
		"bye":               "Bye Bye! See you soon ",
		"good night":        "Goodnight ! Sweet Dreams.",
		"goodnight":         "Goodnight ! Sweet Dreams.",
		"exit":              "type /exit to exit GoBOT",
	}
	response, ok := responses[inf]
	if ok {
		fmt.Println("GoBOT: ", response)
		fmt.Printf("\n")
		mainBot()
	} else {
		if inf == "current time" {
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

		} else if inf == "current date" {
			now := time.Now()
			fmt.Printf("GoBOT> ")
			fmt.Println("Current Date is : ", now.Day(), "/", now.Month(), "/", now.Year())

		} else if inf == "/cls" || inf == "/clear" {
			fmt.Printf("==> Clearing the screen !!!")
			time.Sleep(2 * time.Second)
			main()
		} else if inf == "/exit" {
			fmt.Println("Exiting the GoBOT.....!")
			time.Sleep(1 * time.Second)
			clearScreen()
			os.Exit(1)
		} else if inf == "open setting" || inf == "open settings" {
			fmt.Println("Opening the settings..")
			time.Sleep(1 * time.Second)
			settings()
		} else if inf == "open browser" {
			fmt.Println("wait.....")
			time.Sleep(1 * time.Second)
			openBrowser("brave")
		} else if inf == "open calculator" {
			fmt.Println("Opening the Calculator")
			time.Sleep(1 * time.Second)
			calculator()
		} else if inf == "open notepad" {
			fmt.Println("Wait....")
			time.Sleep(1 * time.Second)
			notePad()
		} else if inf == "open camera" || inf == "open webcam" {
			fmt.Println("Wait....")
			time.Sleep(1 * time.Second)
			openCamera()
		} else {
			fmt.Println("GoBOT> I'm sorry, I don't understand.")
			fmt.Printf("\n")
			fmt.Println("You can use following command such as ")
			info := []string{
				"hello", "open browser", "current time", "current date", "etc",
			}
			for _, disP := range info {
				fmt.Printf("-> ")
				fmt.Println(disP)
			}
		}
		fmt.Printf("\n")
		mainBot()
	}

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
	fmt.Printf("\n")
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
	fmt.Printf("\n")
}
func mediaPlayer() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "start", "wmplayer", "wmplayer.exe", "microphone:")
	cmd.Start()
	fmt.Printf("\n")
}
func settings() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "start", "ms-settings:")
	cmd.Start()
	fmt.Printf("\n")
}
func notePad() {
	cmd := exec.Command("notepad.exe")
	err := cmd.Start()

	if err != nil {
		fmt.Println("Failed to open Notepad..")
	} else {
		fmt.Println("Notepad is opened Sucessfully!")
	}
	fmt.Printf("\n")
}
func openCamera() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "start", "microsoft.windows.camera:")
	cmd.Start()
	fmt.Printf("\n")
}
func main() {
	clearScreen()
	startInterface()
}
