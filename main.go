package main

import (
	"fmt"
	"os"
)

func main() {

	displayIntroduction()
	displayMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring the websites")
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
