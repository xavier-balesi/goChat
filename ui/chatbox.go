package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

type ChatBox struct {
	tview.TextView
}

func (c *ChatBox) AddMessage(message UIMessage) {
	formatMessage := func(u string, m string) string {
		return fmt.Sprintf("%s> %s\n", u, m)
	}
	c.Write([]byte(formatMessage(message.User, message.Message)))
}

func NewChatBox() *ChatBox {
	var tv ChatBox = ChatBox{TextView: *tview.NewTextView()}
	tv.SetTextAlign(tview.AlignLeft).
		SetTextColor(DefaultForegroundColor)
	tv.SetTitle("Chat Box").
		SetBorder(true)
	return &tv
}
