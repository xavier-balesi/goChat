package client

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"simpleChat/config"
	"simpleChat/errors"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

type ClientOptions struct {
	name string
}

var wg sync.WaitGroup

var clientOptions ClientOptions

func main() {
	log.SetLevel(log.DebugLevel)
	user := flag.String("name", "anonynous", "user name for the chat")
	flag.Parse()

	clientOptions = ClientOptions{name: *user}
	log.Debugf("name = %s", clientOptions.name)

	conn, err := net.Dial("tcp", config.Host)
	errors.GestionError(err)

	wg.Add(2)

	go inputHandler(conn)
	go serverHandler(conn)

	wg.Wait()

}

func inputHandler(conn net.Conn) {
	defer wg.Done()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Moi : ")
		text, err := reader.ReadString('\n')
		errors.GestionError(err)

		payload := fmt.Sprintf("Message\t%s\t%s", clientOptions.name, text)
		conn.Write([]byte(payload))
	}
}

func serverHandler(conn net.Conn) {
	defer wg.Done()
	for {
		payload, err := bufio.NewReader(conn).ReadString('\n')
		errors.GestionError(err)

		handlePayload(conn, payload)
		fmt.Print("server : " + payload)
	}
}

func handlePayload(conn net.Conn, payload string) {
	payload = strings.Split(payload, "\n")[0]
	action := strings.Split(payload, "\t")
	log.Debug("\nhandlePayload action = ", action)
	switch action[0] {
	case "Auth":
		log.Debug("incoming auth request")
		conn.Write([]byte(fmt.Sprintf("%s\t%s\n", "auth", clientOptions.name)))
	case "Message":
		log.Debugf("incoming message %s", action[1:])
	default:
		log.Debug("unknown payload")
	}

}
