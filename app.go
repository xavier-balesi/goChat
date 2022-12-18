package main

import (
	"flag"
	"simpleChat/ui"
)

var (
	ch       chan ui.Message
	clientUI *ui.ClientUI
)

func sendHandler(message ui.Message) {
	ch <- message
}

func main() {

	ch = make(chan ui.Message)

	debug_mode := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	clientUI := ui.ClientUI{}
	clientUI.Init(sendHandler, ch, *debug_mode)

	if err := clientUI.Start(); err != nil {
		panic(err)
	}
}
