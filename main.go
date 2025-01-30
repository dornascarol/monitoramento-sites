package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	displayIntroduction()
	readArchiveWebsites()

	for {
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
	fmt.Println("")

	return commandReadyReading
}

func startMonitoring() {
	fmt.Println("Monitoring the websites")

	allWebsites := readArchiveWebsites()

	for i := 0; i < monitoring; i++ {
		for i, website := range allWebsites {
			fmt.Println("Testing website", i, ":", website)
			testWebsite(website)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testWebsite(website string) {
	response, error := http.Get(website)

	if error != nil {
		fmt.Println("An error happened:", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Website:", website, "was loaded successfully!")
	} else {
		fmt.Println("Website:", website, "is having problems. Status code is:", response.StatusCode)
	}
}

func readArchiveWebsites() []string {

	var websites []string

	file, error := os.Open("websites.txt")
	if error != nil {
		fmt.Println("An error happened:", error)
	}

	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)

		if error == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}
