package main

import (
	"fmt"
	"os"
)

// main function
func main() {
	//fmt.Println("Welcome to k8s-doctor")

	ev := CreateNewEventViewer()
	ev.PrintBanner()

	// Set a dummy config
	ct := `k8s-doctor Config comes here`

	tui := NewTUI(ev)
	tui.SetConfig(ct)

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter number of fake events to generate: ")
	//line, err := reader.ReadString('\n')
	//if err != nil {
	//	ev.PrintError("Failed to read input")
	//	os.Exit(1)
	//}
	//
	//line = strings.TrimSpace(line)
	//n, err := strconv.Atoi(line)
	//if err != nil {
	//	ev.PrintError("Invalid input")
	//	os.Exit(1)
	//}

	n := 6 // default number of events
	events := GenerateFakeEvents(n)
	ev.PrintInfo(fmt.Sprintf("Generated %d fake events", len(events)))
	fmt.Println()

	//ev.PrintAllEvents(events)
	tui.DisplayEvents(events)

	if err := tui.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
