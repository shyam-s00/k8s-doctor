package main

import (
	"io"

	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	app          *tview.Application
	configView   *tview.TextView
	outputView   *tview.TextView
	inputField   *tview.InputField
	eventViewer  *EventViewer
	layout       *tview.Flex
	outputWriter io.Writer
}

func NewTUI(ev *EventViewer) *TUI {
	tui := &TUI{
		app:         tview.NewApplication(),
		configView:  tview.NewTextView(),
		outputView:  tview.NewTextView(),
		inputField:  tview.NewInputField(),
		eventViewer: ev,
	}

	//upper half of the screen
	tui.configView.SetTitle(" Configuration ").
		SetBorder(true)

	// second half or lower half of the screen
	tui.outputView.SetTitle(" Events ").
		SetBorder(true)

	// forward the output to the outputView
	tui.outputWriter = tview.ANSIWriter(tui.outputView)
	color.Output = tui.outputWriter

	tui.inputField.SetLabel("> ").
		SetDoneFunc(func(key tcell.Key) {}) // TODO: Add input handling

	tui.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tui.configView, 0, 1, false).
		AddItem(tui.outputView, 0, 1, false).
		AddItem(tui.inputField, 1, 0, true)

	return tui
}

func (tui *TUI) SetConfig(config string) {
	tui.configView.SetText(config)
}

func (tui *TUI) ShowEvent(e Event) {
	tui.eventViewer.DisplayEvent(e)
	tui.outputView.ScrollToEnd()
}

func (tui *TUI) DisplayEvents(events []Event) {
	for _, e := range events {
		tui.ShowEvent(e)
	}
}

func (tui *TUI) Run() error {
	return tui.app.SetRoot(tui.layout, true).Run()
}
