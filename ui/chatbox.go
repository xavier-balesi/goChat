package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

type ChatBox struct {
	tview.TextView
}

func (c *ChatBox) AddMessage(user string, message string) {
	formatMessage := func(u string, m string) string {
		return fmt.Sprintf("%s> %s\n", u, m)
	}
	c.Write([]byte(formatMessage(user, message)))
}

func NewChatBox() *ChatBox {
	var tv ChatBox = ChatBox{TextView: *tview.NewTextView()}
	tv.SetTextAlign(tview.AlignLeft).
		SetTextColor(DefaultForegroundColor)
	tv.SetTitle("Chat Box").
		SetBorder(true)
	return &tv
}
