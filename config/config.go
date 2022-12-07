package config

import "fmt"

const (
	DEFAULT_IP   = "127.0.0.1"
	DEFAULT_PORT = "3569"
)

var Host = fmt.Sprintf(
	"%s:%s",
	DEFAULT_IP,
	DEFAULT_PORT,
)
