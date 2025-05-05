# Glô Monitoring :mag_right:

_Project feature I learned at Alura_

## Introduction
This feature is responsible for monitoring the availability of selected Glô websites, verifying whether they are online or offline. All monitoring data, including the exact date, time, URL, and availability status, is systematically recorded in a log file.


## Technologies used
* VS Code
* Go (Golang)
* Package fmt
* Package net/http
* Package os
* Package time
* Package bufio
* Package io
* Package strconv
* Package strings


## Tools
The <a href="https://go.dev/doc/install" target="_blank"> Go</a> programming language was installed on Windows, version 1.22.2

The Go extension was installed in <a href="https://code.visualstudio.com/download" target="_blank"> VS Code</a>, version 0.44.0

Predefined layouts were used with <a href="https://go.dev/src/time/format.go" target="_blank"> Time.Format</a> for custom date and time specific formatting.


## Running the project
1. **To get the repository, clone it and go to the project folder:**
	```bash
	git clone https://github.com/dornascarol/monitoramento-sites.git
	cd monitoramento-sites
	```

2. **Command to run the project in the terminal:**
	```
	go run main.go
	```

3. **To stop the execution in the terminal: press `"Ctrl"` + `"C"`.**
<br>

### _Menu in terminal_
- Selecting Option 1: starts monitoring all websites listed in `websites.txt`. At the same time, the `logs.txt` file will record all logs, including full date, time, URL, and status.

- Selecting Option 2: prints the logs. The logs are initially returned as a byte array but are converted to a string for display in the terminal.

- Selecting Option 3: exits the program. 


## Project status
:heavy_check_mark: Application completed.
