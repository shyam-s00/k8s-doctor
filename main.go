package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main function
func main() {
	//fmt.Println("Welcome to k8s-doctor")
	ev := CreateNewEventViewer()
	ev.PrintBanner()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of fake events to generate: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		ev.PrintError("Failed to read input")
		os.Exit(1)
	}

	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil {
		ev.PrintError("Invalid input")
		os.Exit(1)
	}
	events := GenerateFakeEvents(n)
	ev.PrintInfo(fmt.Sprintf("Generated %d fake events", len(events)))
	fmt.Println()
	ev.PrintAllEvents(events)
}
