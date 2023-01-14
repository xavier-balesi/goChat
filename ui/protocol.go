package ui

type UIProtocolAdapter struct {
	incomingMessageChannel chan UIMessage
	outgoingMessageChannel chan string

	clientUI ClientUI
}

func (c *UIProtocolAdapter) Init(incomingChan chan UIMessage, outgoingChan chan string, clientUI ClientUI) {
	c.incomingMessageChannel = incomingChan
	c.outgoingMessageChannel = outgoingChan
	c.clientUI = clientUI
}

func (c *UIProtocolAdapter) onNewMessage(message string) {

}

func (c *UIProtocolAdapter) onSendMessage(message string) {
	c.outgoingMessageChannel <- message
}
