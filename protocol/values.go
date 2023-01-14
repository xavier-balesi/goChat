package protocol

const PrefixMessage string = "Message"
const PrefixAuth string = "Auth"

type Message struct {
	User    string
	Message string
}
