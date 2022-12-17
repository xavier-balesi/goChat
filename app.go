package main

import (
	"flag"
	"simpleChat/ui"
)

var (
	ch       chan string
	clientUI *ui.ClientUI
)

func sendHandler(sentence string) {
	ch <- sentence
}

func main() {

	ch = make(chan string)

	debug_mode := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	clientUI := ui.ClientUI{}
	clientUI.Init(sendHandler, ch, *debug_mode)

	if err := clientUI.Start(); err != nil {
		panic(err)
	}
}
