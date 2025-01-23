package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	displayIntroduction()
	displayMenu()

	command := readCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Displaying the logs")
	case 3:
		fmt.Println("Exiting the program")
		os.Exit(0)
	default:
		fmt.Println("I don't know this command")
		os.Exit(-1)
	}
}

func displayIntroduction() {
	name := "Glô"
	version := 1.1
	fmt.Println("Hi,", name)
	fmt.Println("This program is in version", version)
}

func displayMenu() {
	fmt.Println("1) Start websites monitoring")
	fmt.Println("2) Display the logs")
	fmt.Println("3) Exit the program")
}

func readCommand() int {
	var commandReadyReading int
	fmt.Scan(&commandReadyReading)
	fmt.Println("The command selected was:", commandReadyReading)

	return commandReadyReading
}

func startMonitoring() {
	fmt.Println("Monitoring the websites")
	website := "https://ge.globo.com/"
	response, _ := http.Get(website)

	if response.StatusCode == 200 {
		fmt.Println("Website:", website, "was loaded successfully!")
	} else {
		fmt.Println("Website:", website, "is having problems. Status code is:", response.StatusCode)
	}
}
