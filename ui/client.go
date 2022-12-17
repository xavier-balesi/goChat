package ui

import (
	"time"

	"github.com/rivo/tview"
)

const refreshInterval = 100 * time.Millisecond

type ClientUI struct {
	chatBox   *ChatBox
	logBox    *LogBox
	sendInput *tview.InputField
	app       *tview.Application
	grid      *tview.Grid
	ch        chan string
}

func (c *ClientUI) Init(sendHandler func(string), ch chan string, debugNode bool) {

	c.chatBox = NewChatBox()
	c.sendInput = NewSendBox(sendHandler)
	c.logBox = NewLogBox()
	c.ch = ch

	columns := []int{80}
	if debugNode {
		columns = append(columns, 80)
	}

	c.grid = tview.NewGrid().
		SetRows(10, 1).
		SetColumns(columns...).
		SetBorders(false).
		AddItem(c.chatBox, 0, 0, 1, 1, 0, 0, false).
		AddItem(c.sendInput, 1, 0, 1, 1, 0, 0, true)

	if debugNode {
		c.grid.AddItem(c.logBox, 0, 1, 2, 1, 0, 0, false)
	}

	c.app = tview.NewApplication()
}

func (c *ClientUI) channelListener() {
	for {
		time.Sleep(refreshInterval)
		c.app.QueueUpdateDraw(func() {
			select {
			case sentence := <-c.ch:
				c.logBox.AddLog("DEBUG", "add sentence "+sentence+"\n")
				c.chatBox.AddMessage("moi", sentence)
			default:
			}
		})
	}
}

func (c *ClientUI) Start() error {
	go c.channelListener()
	return c.app.SetRoot(c.grid, true).Run()
}
