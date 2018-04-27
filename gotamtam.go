package gotamtam

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	Token      string
	Version    string
	Name       string
	Seq        int
}

// GoTamTamReader interface for reader
type Reader interface {
	Response(client *Client, message *Message)
}

// Create NEW client
func New(token, version, name string) (*Client, error) {
	tamtam := &Client{}
	u := url.URL{Scheme: "wss", Host: "tamtam-ws.ok.ru", Path: "/websocket"}

	d := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	h := http.Header{}

	c, _, err := d.Dial(u.String(), h)
	if err != nil {
		return nil, err
	}

	tamtam.Connection = c
	tamtam.Token = token
	tamtam.Version = version
	tamtam.Name = name
	tamtam.Seq = 1

	return tamtam, nil
}

func (c *Client) Serve(reader Reader) {
	go func() {
		for {
			_, bytes, err := c.Connection.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
			}
			message, err := DecomposeMessage(bytes)
			if err != nil {
				log.Println("Decompose error:", err)
			}
			reader.Response(c, message)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	login := time.NewTimer(time.Second)
	ping := time.NewTicker(time.Second * 30)
	done := make(chan struct{})

	for {
		select {
		case <-done:
			return
		case <-ping.C:
			c.Write(PING, InteractivePayload{true})
		case <-login.C:
			var UserAgent = UserAgent{
				DeviceType: "BOT",
				AppVersion: "0.0.1",
				DeviceName: "TestBOT",
			}

			c.Write(SESSION_INIT, HelloPayload{UserAgent: UserAgent})
			c.Write(PING, InteractivePayload{true})
			c.Write(LOGIN, LoginPayload{
				ChatsSync:    makeTimestamp(),
				ConfigHash:   "59256f37-0000000000000000-80000020-00000162ca55d8f3-f97f7333-0000000000000000-2e30ed1c-2",
				ContactSync:  makeTimestamp(),
				PresenceSync: makeTimestampSeconds(),
				Token:        c.Token,
				UserAgent:    UserAgent,
			})
		case <-interrupt:
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.Connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second * 10):
			}
			return
		}
	}
}

func (c *Client) SendMessage(chatID int, chatType, text string) {
	message := SendMessagePayload{
		ChatID: chatID,
		Message: SendMessage{
			Text:        text,
			CID:         makeTimestamp(),
			DetectShare: true,
		},
		Notify: true,
		Type:   chatType,
	}
	c.Write(MSG_SEND, message)
}

func (c *Client) Write(opCode int, payload interface{}) {
	b, err := ComposeMessage(c.Seq, opCode, payload)
	if err != nil {
		log.Println("Compose error:", err)
		return
	}
	fmt.Println(string(b))
	err = c.Connection.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		log.Println("Write error:", err)
		return
	}
	c.Seq++
}

// ComposeMessage compose json byte message
func ComposeMessage(seq int, opCode int, payload interface{}) ([]byte, error) {
	b, err := json.Marshal(
		Message{
			Ver:     10,
			Cmd:     0,
			Seq:     seq,
			OpCode:  opCode,
			Payload: payload,
		},
	)
	return b, err
}

func DecomposeMessage(data []byte) (*Message, error) {
	r := &Message{}
	err := json.Unmarshal(data, r)
	if err != nil {
		return nil, err
	}

	switch r.OpCode {
	case NOTIF_MESSAGE:
		var result NotifyMessagePayload
		err := mapstructure.Decode(r.Payload, &result)
		if err != nil {
			return nil, err
		}
		r.Payload = result
	}
	return r, err
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func makeTimestampSeconds() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}
