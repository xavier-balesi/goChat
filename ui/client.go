package ui

import (
	"time"

	"github.com/rivo/tview"
)

const refreshInterval = 100 * time.Millisecond

type Message struct {
	user    string
	message string
}

type ClientUI struct {
	chatBox   *ChatBox
	logBox    *LogBox
	sendInput *tview.InputField
	app       *tview.Application
	flexGrid  *tview.Flex
	incomCh   chan Message
}

func (c *ClientUI) createGrid(debugMode bool) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(c.chatBox, 0, 4, false).
			AddItem(c.sendInput, 1, 1, true), 0, 2, true)

	if debugMode {
		flex.AddItem(c.logBox, 0, 1, false)
	}

	return flex
}

func (c *ClientUI) Init(sendHandler func(Message), ch chan Message, debugMode bool) {

	c.chatBox = NewChatBox()
	c.sendInput = NewSendBox(sendHandler)
	c.logBox = NewLogBox()
	c.flexGrid = c.createGrid(debugMode)
	c.incomCh = ch

	c.app = tview.NewApplication()
}

func (c *ClientUI) channelListener() {
	for {
		time.Sleep(refreshInterval)
		c.app.QueueUpdateDraw(func() {
			select {
			case sentence := <-c.incomCh:
				c.logBox.AddLog("DEBUG", "add sentence "+sentence.user+":"+sentence.message+"\n")
				c.chatBox.AddMessage(sentence.user, sentence.message)
			default:
			}
		})
	}
}

func (c *ClientUI) Start() error {
	go c.channelListener()
	return c.app.SetRoot(c.flexGrid, true).Run()
}
