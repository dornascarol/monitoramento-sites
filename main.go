package main

import "fmt"

func main() {
	name := "Glô"
	version := 1.0

	fmt.Println("Hi,", name)
	fmt.Println("This program is in version", version)

	fmt.Println("1) Start websites monitoring")
	fmt.Println("2) Display the logs")
	fmt.Println("3) Exit the program")

	var command int
	fmt.Scan(&command)
	fmt.Println("The command selected was:", command)

	switch command {
	case 1:
		fmt.Println("Monitoring the websites")
	case 2:
		fmt.Println("Displaying the logs")
	case 3:
		fmt.Println("Exiting the program")
	default:
		fmt.Println("I don't know this command")
	}
}
