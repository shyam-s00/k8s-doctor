package main

import (
	"sort"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// EventType Simple event model used by the event viewer
// TODO: Move event model to separate folders
type EventType string

const (
	EventTypeInfo    EventType = "INFO"
	EventTypeError   EventType = "ERROR"
	EventTypeWarning EventType = "WARNING"
	EventTypeDebug   EventType = "DEBUG"
)

type Event struct {
	Type      EventType
	Timestamp time.Time
	Icon      string
	Name      string
	Message   string
}

type EventViewer struct {

	// Color definitions
	info    *color.Color
	error   *color.Color
	warning *color.Color
	debug   *color.Color
	banner  *color.Color
	// success and failures
	success *color.Color
	failure *color.Color
}

func CreateNewEventViewer() *EventViewer {
	return &EventViewer{
		banner:  color.New(color.FgCyan, color.Bold),
		info:    color.New(color.FgBlue),
		warning: color.New(color.FgYellow),
		error:   color.New(color.FgRed, color.Bold),
		debug:   color.New(color.FgHiBlack),
		success: color.New(color.FgGreen),
		failure: color.New(color.FgHiRed),
	}
}

func (p *EventViewer) PrintBanner() {
	b := ` 🩺  k8s-doctor`
	p.banner.Println(b)
}

// PrintInfo prints an informational message.
func (p *EventViewer) PrintInfo(msg string) {
	p.info.Printf("ℹ %s\n", msg)
}

// PrintSuccess prints a success message.
func (p *EventViewer) PrintSuccess(msg string) {
	p.success.Printf("✓ %s\n", msg)
}

// PrintWarning prints a warning message.
func (p *EventViewer) PrintWarning(msg string) {
	p.warning.Printf("⚠ %s\n", msg)
}

// PrintError prints an error message.
func (p *EventViewer) PrintError(msg string) {
	p.error.Printf("✖ %s\n", msg)
}

func (p *EventViewer) DisplayEvent(e Event) {
	timestamp := e.Timestamp.Format("2006-01-02 15:04:05")
	switch e.Type {
	case EventTypeInfo:
		p.info.Printf("[%s] %s %-8s %s - %s\n", timestamp, e.Icon, string(e.Type), e.Name, e.Message)
	case EventTypeDebug:
		p.debug.Printf("[%s] %s %-8s %s - %s\n", timestamp, e.Icon, string(e.Type), e.Name, e.Message)
	case EventTypeError:
		p.error.Printf("[%s] %s %-8s %s - %s\n", timestamp, e.Icon, string(e.Type), e.Name, e.Message)
	case EventTypeWarning:
		p.warning.Printf("[%s] %s %-8s %s - %s\n", timestamp, e.Icon, string(e.Type), e.Name, e.Message)
	default:
		p.info.Printf("[%s] %s %-8s %s - %s\n", timestamp, e.Icon, string(e.Type), e.Name, e.Message)
	}
}

func (p *EventViewer) PrintAllEvents(events []Event) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Before(events[j].Timestamp)
	})

	// TODO: Modify this to support continuous output
	p.debug.Println("--- Event Log Start ---")
	for _, e := range events {
		p.DisplayEvent(e)
	}

	p.debug.Println("--- Event Log End ---")
}

// GenerateFakeEvents Stub for generating fake events for now
func GenerateFakeEvents(n int) []Event {
	events := make([]Event, 0, n)
	now := time.Now()

	typeset := []EventType{EventTypeInfo, EventTypeWarning, EventTypeError, EventTypeDebug}
	icons := map[EventType]string{
		EventTypeInfo: "🔔",
		//EventTypeNotice: "🚀",
		EventTypeWarning: "⚠️",
		EventTypeError:   "💥",
		EventTypeDebug:   "🐞",
	}

	messages := []string{
		"Deployment updated",
		"Pod restarted",
		"Image rolled out",
		"OOMKilled detected",
		"Manual rollout restart",
		"Autoscaler adjusted replicas",
		"Liveness probe failed",
	}

	for i := 0; i < n; i++ {
		events = append(events, Event{
			Type:      typeset[i%len(typeset)],
			Timestamp: now.Add(time.Duration(i) * time.Second),
			Icon:      icons[typeset[i%len(typeset)]],
			Name:      "EventName" + strconv.Itoa(i+1),
			Message:   messages[i%len(messages)],
		})
	}

	return events
}
