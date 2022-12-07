package main

import (
	"bufio"
	"errors"
	"net"
	"simpleChat/config"
	chatErrors "simpleChat/errors"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ClientInfo struct {
    name string
    conn net.Conn
}

var clients []ClientInfo

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Starting chat server...")

	// listen
	ln, err := net.Listen("tcp", config.Host)
	chatErrors.GestionError(err)

	for {
		// accept connexions
		conn, err := ln.Accept()
		chatErrors.GestionError(err)
		log.Info("new cient connection from ", conn.RemoteAddr())

		clientName, authErr := authentify(conn)
		if authErr != nil {
			log.Errorf("client not authentified : %s", conn.RemoteAddr())
			continue
		}

        clients = append(clients, ClientInfo{name: clientName, conn: conn})
		go clientHandler(conn)
	}
}

func authentify(conn net.Conn) (string, error) {
	var authRequest string = "Auth\n"
	conn.Write([]byte(authRequest))

	buf := bufio.NewReader(conn)
	payload, err := buf.ReadString('\n')
	if err != nil {
		log.Warn("Client disconnected :", conn.RemoteAddr())
		return "", errors.New("cannot read client")
	}

	payload = strings.Split(payload, "\n")[0]

	log.Debugf("received payload : %s", payload)
	action, clientName := splitPayload(payload)
	if action == "auth" && isAuthorized(clientName) {
        log.Infof("client authentified : %s / %s", conn.RemoteAddr(), clientName)
		return clientName, nil
	}

	return "", errors.New("Wrong auth !")
}

func isAuthorized(name string) bool {
    var Authorized = map[string]bool{
        "xavier": true,
        "toto": false,
    }
    v, exists := Authorized[name]
    if exists { return v }
    return false
}

func splitPayload(payload string) (string, string) {
	splitted := strings.Split(payload, "\t")
	return splitted[0], splitted[1]
}

func clientHandler(conn net.Conn) {
	buf := bufio.NewReader(conn)
	for {
		payload, err := buf.ReadString('\n')
		if err != nil {
			log.Info("Client disconnected :", conn.RemoteAddr())
			break
		}

		payload = strings.Split(payload, "\n")[0]
		log.Debugf("[%s] =msg=> %s", conn.RemoteAddr(), payload)

        response, err := handlePayload(payload)

		for _, c := range clients {
            log.Debugf("writing to client %s : '%s'", c.name, response)
			c.conn.Write([]byte(response))
		}
	}
}

func handlePayload(payload string) (string, error) {
	payload_protocol := strings.Split(payload, "\t")
	log.Debugf("payload_protocol=%s", payload_protocol)
    response := []string{"Message", payload_protocol[1], payload_protocol[2]}

    return strings.Join(response[:], "\t") + "\n", nil
}
