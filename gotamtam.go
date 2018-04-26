package gotamtam

import "encoding/json"

// ComposeMessage compose json byte message
func ComposeMessage(seq int, opCode int, payload interface{}) ([]byte, error) {
	b, err := json.Marshal(
		Wrapper{
			Ver:     10,
			Cmd:     0,
			Seq:     seq,
			OpCode:  opCode,
			Payload: payload,
		},
	)
	return b, err
}
