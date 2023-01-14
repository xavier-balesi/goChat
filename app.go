package main

import (
	"flag"
	"os"
	"simpleChat/config"
	"simpleChat/protocol"
	"simpleChat/server"
	"simpleChat/ui"

	log "github.com/sirupsen/logrus"
)

var (
	messageCh chan ui.UIMessage
	clientUI  *ui.ClientUI
)

func makeSendHandler(user string) func(string) {
	return func(message string) {
		mes := ui.UIMessage{
			User:    user,
			Message: message,
		}
		messageCh <- mes
	}
}

type UIProtocolHandler struct {
	messageCh chan ui.UIMessage
}

func (c *UIProtocolHandler) Init(msgCh chan ui.UIMessage) {
	c.messageCh = msgCh
}
func (c *UIProtocolHandler) OnNewMessage(message protocol.Message) {
	c.messageCh <- ui.UIMessage{User: message.User, Message: message.Message}
}

func startClient(userName string, debugMode bool) {

	logFile, _ := os.OpenFile("client.log", os.O_WRONLY|os.O_CREATE, 0644)
	log.SetOutput(logFile)

	var serverConfig server.ServerConfig
	serverConfig.Host = config.DEFAULT_IP
	serverConfig.Port = config.DEFAULT_PORT

	var uiProtocolHandler = new(UIProtocolHandler)
	uiProtocolHandler.Init(messageCh)

	server := new(server.BaseServerHandler)
	server.Init(serverConfig, uiProtocolHandler)
	server.Connect()

	clientUI := ui.ClientUI{}
	clientUI.Init(server.SendMessage, messageCh, debugMode)

	if err := clientUI.Start(); err != nil {
		panic(err)
	}
}

func startServer() {
	server.Main()
}

func main() {
	messageCh = make(chan ui.UIMessage)

	debugMode := flag.Bool("debug", false, "enable debug mode")
	userName := flag.String("user", "", "user name")
	startType := flag.String("type", "server", "client | server")
	flag.Parse()

	if *debugMode {
		log.SetLevel(log.DebugLevel)
	}
	if *startType == "server" {
		startServer()
	} else {
		startClient(*userName, *debugMode)
	}

}
