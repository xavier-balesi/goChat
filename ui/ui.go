package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const refreshInterval = 100 * time.Millisecond
const foregroundColor = tcell.ColorWheat

var (
	app       *tview.Application
	chatView  *tview.TextView
	logText   *tview.TextView
	sendInput *tview.InputField
	ch        chan string
)

func addLog(level string, message string) {
	logMessage := fmt.Sprintf("[%s] %s", level, message)
	logText.Write([]byte(logMessage))
}

func makeInputHandler(input *tview.InputField) func(tcell.Key) {
	return func(key tcell.Key) {
		if key == tcell.KeyEnter {
			go func() {
				sentence := input.GetText()
				ch <- sentence
				input.SetText("")
			}()
		}
	}
}

func channelListener() {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
			select {
			case sentence := <-ch:
				addLog("DEBUG", "add sentence "+sentence+"\n")
				fullText := chatView.GetText(false)
				fullText += sentence
				chatView.SetText(fullText)
			default:
			}
		})
	}
}

func createChatBox() *tview.TextView {
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetText("hello\nWorld").
		SetTextColor(foregroundColor)
	tv.SetTitle("Chat Box").
		SetBorder(true)
	return tv
}

func main() {
	// log.SetLevel(log.DebugLevel)

	ch = make(chan string)

	debug_mode := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	chatView = createChatBox()

	logText = tview.NewTextView().SetTextAlign(tview.AlignLeft)
	logText.SetTitle("Log Box").SetBorder(true)

	sendInput = tview.NewInputField()
	sendInput.SetDoneFunc(makeInputHandler(sendInput))
	sendInput.SetFieldBackgroundColor(tcell.ColorBlack)
	sendInput.SetFieldTextColor(foregroundColor)
	sendInput.SetLabel("> ")

	columns := []int{80}
	if *debug_mode {
		columns = append(columns, 80)
	}

	grid := tview.NewGrid().
		SetRows(10, 1).
		SetColumns(columns...).
		SetBorders(false).
		AddItem(chatView, 0, 0, 1, 1, 0, 0, false).
		AddItem(sendInput, 1, 0, 1, 1, 0, 0, true)

	if *debug_mode {
		grid.AddItem(logText, 0, 1, 2, 1, 0, 0, false)
	}

	go channelListener()

	app = tview.NewApplication()
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
