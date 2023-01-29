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

var cmd *exec.Cmd

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
	fmt.Printf("\n")

	mainBot()
}

func mainBot() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("==> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	if input == "/help" {
		fmt.Printf("\n")
		goHelp()
	} else if input == "/list" {
		fmt.Printf("\n")
		goList()
	}
	time.Sleep(1 * time.Second)
	reply(input)

}
func goList() {
	fmt.Println("GoBOT> Here are the list of instruction you can use!")
	lists := []string{
		"open browser", "open camera", "open settings", "current date", "current time", "open notepad", "open explorer",
		"open mediaplayer", "open paint", "open calculator", "hi", "hello", "good morning", "bye", "good night",
	}

	for _, list := range lists {
		fmt.Printf("-> ")
		fmt.Println(list)
	}
	fmt.Printf("\n")
	mainBot()
}
func goHelp() {
	fmt.Println("GoBOT> Here are the command you can use in GoBOT!")

	cmdList := []string{
		"to see the set the set of instructuions type /list ", "to clear the history use  /cls or /clear", "to exit GoBOT use /exit", "to shutdown computer use /shutdown",
	}
	for _, cmdL := range cmdList {
		fmt.Printf("-> ")
		fmt.Println(cmdL)
	}
	fmt.Printf("\n")
	mainBot()
}
func reply(sentence string) {
	inf := sentence
	responses := map[string]string{
		"hi":                "Hi! how are you ?",
		"hello":             "Hello !",
		"i am fine and you": "I'm good, thank you! How may i help you?",
		"what is your name": "My name is GoBOT!",
		"how are you":       "I'm good, thank you!",
		"good morning":      "Goodmorning ! Have a good Day!",
		"goodmorning":       "Goodmorning ! Have a good Day!",
		"bye":               "Bye Bye! See you soon ",
		"bye bye":           "Bye Bye! See you soon ",
		"good night":        "Goodnight ! Sweet Dreams.",
		"goodnight":         "Goodnight ! Sweet Dreams.",
		"exit":              "type /exit to exit GoBOT",
		"clear":             "type /clear to clear the results",
		"shutdown":          "type /shutdown to shudown computer",
		"list":              "type /list to see the instructions",
		"help":              "type /help if you need any Help",
	}
	response, ok := responses[inf]
	if ok {
		fmt.Printf("\n")
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
			fmt.Printf("\n")
			now := time.Now()
			fmt.Printf("GoBOT> ")
			fmt.Println("Current Date is : ", now.Day(), "/", now.Month(), "/", now.Year())

		} else if inf == "/cls" || inf == "/clear" {
			fmt.Printf("\n")
			fmt.Printf("==> Clearing the screen !!!")
			time.Sleep(1 * time.Second)
			main()
		} else if inf == "/exit" {
			fmt.Printf("\n")
			fmt.Println("Exiting the GoBOT.....!")
			time.Sleep(1 * time.Second)
			clearScreen()
			os.Exit(1)
		} else if inf == "open setting" || inf == "open settings" {
			fmt.Printf("\n")
			fmt.Println("Opening the settings....")
			time.Sleep(1 * time.Second)

			settings()
		} else if inf == "open browser" {
			fmt.Printf("\n")
			fmt.Println("Opening the browser....")
			time.Sleep(1 * time.Second)
			openBrowser()
		} else if inf == "open calculator" {
			fmt.Printf("\n")
			fmt.Println("Opening the Calculator....")
			time.Sleep(1 * time.Second)
			calculator()
		} else if inf == "open notepad" {
			fmt.Printf("\n")
			fmt.Println("Opening the notepad....")
			time.Sleep(1 * time.Second)
			notePad()
		} else if inf == "open camera" || inf == "open webcam" {
			fmt.Printf("\n")
			fmt.Println("Opening the camera....")
			time.Sleep(1 * time.Second)
			openCamera()
		} else if inf == "open explorer" {
			fmt.Printf("\n")
			fmt.Println("Opening the file explorer....")
			time.Sleep(1 * time.Second)
			openExplorer()
		} else if inf == "open mediaplayer" {
			fmt.Printf("\n")
			fmt.Println("Opening the media player....")
			time.Sleep(1 * time.Second)
			mediaPlayer()
		} else if inf == "/shutdown" {
			fmt.Printf("\n")
			fmt.Println("Wait....")
			time.Sleep(1 * time.Second)
			shutDown()
		} else if inf == "open paint" {
			fmt.Printf("\n")
			fmt.Println("Opening the paint....")
			time.Sleep(1 * time.Second)
			openPaint()
		} else {
			fmt.Printf("\n")
			fmt.Println("GoBOT> I'm sorry, I don't understand.")
			fmt.Printf("\n")
			fmt.Println("You can use following command such as ")
			info := []string{
				"hello", "open browser", "current time", "current date", "open mediaplayer", "etc",
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
func openPaint() {
	cmd := exec.Command("cmd", "/c", "start", "mspaint")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open paint!")
	} else {
		fmt.Println("Paint opened sucessfully!")
	}
}
func shutDown() {
	cmd := exec.Command("shutdown", "-h", "now")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func openExplorer() {
	cmd = exec.Command("cmd", "/c", "start", "explorer.exe")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open File Explorer")
	} else {
		fmt.Println("File explorer opened sucessfully!")
	}

}
func openBrowser() {
	cmd = exec.Command("cmd", "/c", "start", "brave.exe")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open Browser")
	} else {
		fmt.Println("Browser opened sucessfully!")
	}
}

func calculator() {
	cmd := exec.Command("calc.exe")

	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open calculator")
	} else {
		fmt.Println("Calculator opened sucessfully!")
	}

}
func mediaPlayer() {
	cmd = exec.Command("cmd", "/C", "start", "wmplayer.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to open MediaPlayer")
	} else {
		fmt.Println("Mediaplayer opened sucessfully!")
	}
}
func settings() {
	cmd = exec.Command("cmd", "/C", "start", "ms-settings:")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to open Settings!")
	} else {
		fmt.Println("Settings opened sucessfully!")
	}
}
func notePad() {
	cmd := exec.Command("notepad.exe")
	err := cmd.Start()

	if err != nil {
		fmt.Println("Failed to open Notepad!")
	} else {
		fmt.Println("Notepad is opened Sucessfully!")
	}
}
func openCamera() {
	cmd = exec.Command("cmd", "/c", "start", "microsoft.windows.camera:")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to open Camera!")
	} else {
		fmt.Println("Camera opened sucessfully!")
	}

}
func main() {
	clearScreen()
	startInterface()
}
