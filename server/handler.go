package server

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	chaterrors "simpleChat/errors"
	"simpleChat/protocol"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ProtocolHandler interface {
	OnNewMessage(protocol.Message)
}

type ServerConfig struct {
	Host string
	Port string
}

type BaseServerHandler struct {
	config          ServerConfig
	conn            net.Conn
	protocolHandler ProtocolHandler
}

func (s *BaseServerHandler) Init(config ServerConfig, protocolHandler ProtocolHandler) {
	s.config = config
	s.protocolHandler = protocolHandler
}

func (s *BaseServerHandler) payloadListener() {
	for {
		payload, err := bufio.NewReader(s.conn).ReadString('\n')
		chaterrors.ErrorHandler(err)
		log.Debugf("received payload from server [%s]", payload)

		err = s.handlePayload(payload)
		chaterrors.ErrorHandler(err)
	}
}

func (s *BaseServerHandler) handlePayload(payload string) error {
	payload = strings.Split(payload, "\n")[0]
	splitted_payload := strings.Split(payload, "\t")
	log.Debug("splitted_payload = ", splitted_payload)

	action := splitted_payload[0]
	switch action {
	case protocol.PrefixMessage:
		s.protocolHandler.OnNewMessage(
			protocol.Message{
				User:    splitted_payload[1],
				Message: splitted_payload[2]})
	case protocol.PrefixAuth:
		s.conn.Write([]byte(fmt.Sprintf("%s\t%s\n", "auth", "xavier")))
	default:
		return errors.New(fmt.Sprintf("Unknown protocol action [%s/%s]", action, splitted_payload[1:]))
	}
	return nil
}

func (s *BaseServerHandler) onNewMessage(message string) {
	log.Debugf("incoming Message [%s]", message)
}

func (s *BaseServerHandler) Connect() {
	var url string = fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	log.Debug("connecting to server at ", url)

	conn, err := net.Dial("tcp", url)
	chaterrors.ErrorHandler(err)

	log.Info("Connected to server at ", url)
	s.conn = conn

	go s.payloadListener()
}

func (s *BaseServerHandler) SendMessage(message string) {
	payload := fmt.Sprintf("%s\t%s\n", protocol.PrefixMessage, message)
	log.Debugf("sending payload [%s]", payload)
	s.conn.Write([]byte(payload))
}
