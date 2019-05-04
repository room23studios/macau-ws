package proto

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// ParseMessage takes the bytes sent in a websocket message and returns a
// pointer to the unmarshaled command as an `interface{}`. You can then use a
// type switch to handle the command.
func ParseMessage(data []byte) (interface{}, error) {
	var m Message
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	var command interface{}

	switch m.Command {
	case "hello":
		command = &CommandHello{}
	case "ping":
		command = &CommandPing{}
	default:
		return nil, fmt.Errorf("unknown command")
	}

	err = json.Unmarshal(m.Data, command)
	if err != nil {
		return nil, err
	}

	return command, err
}
