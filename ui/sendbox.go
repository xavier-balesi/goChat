package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func makeInputHandler(input *tview.InputField, inputHandler func(string)) func(tcell.Key) {
	return func(key tcell.Key) {
		if key == tcell.KeyEnter {
			sentence := input.GetText()
			go inputHandler(sentence)
			input.SetText("")
		}
	}
}

func NewSendBox(inputHandler func(string)) *tview.InputField {
	var sb *tview.InputField = tview.NewInputField()
	sb.SetDoneFunc(makeInputHandler(sb, inputHandler))
	sb.SetFieldBackgroundColor(tcell.ColorBlack)
	sb.SetFieldTextColor(DefaultForegroundColor)
	sb.SetLabel("> ")
	return sb
}
