package ui

import (
	"fmt"
	"github.com/rivo/tview"
)

type LogBox struct {
    tview.TextView
}

func (l *LogBox) AddLog(level string, message string) {
    formatLog := func (l string, m string) string {
        return fmt.Sprintf("[%s] %s", l, m)
    }
	l.Write([]byte(formatLog(level, message)))
}

func NewLogBox() *LogBox{
    var lb LogBox= LogBox{TextView: *tview.NewTextView()}
	lb.SetTextAlign(tview.AlignLeft)
	lb.SetTitle("Log Box").SetBorder(true)
    return &lb
}

