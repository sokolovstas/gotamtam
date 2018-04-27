package gotamtam

type Message struct {
	Ver      int         `json:"ver"`
	Cmd      int         `json:"cmd"`
	Seq      int         `json:"seq"`
	OpCode   int         `json:"opcode"`
	Payload  interface{} `json:"payload"`
	Duration int64       `json:"duration,omitempty"`
}

type HelloPayload struct {
	UserAgent UserAgent `json:"userAgent"`
}

type UserAgent struct {
	DeviceType      string `json:"deviceType"`
	AppVersion      string `json:"appVersion"`
	OsVersion       string `json:"osVersion"`
	Locale          string `json:"locale"`
	DeviceName      string `json:"deviceName"`
	Screen          string `json:"screen"`
	HeaderUserAgent string `json:"headerUserAgent"`
}

type InteractivePayload struct {
	Interactive bool `json:"interactive"`
}

type LoginPayload struct {
	ChatsSync    int64     `json:"chatsSync"`
	ConfigHash   string    `json:"configHash"`
	ContactSync  int64     `json:"contactsSync"`
	PresenceSync int64     `json:"presenceSync"`
	Token        string    `json:"token"`
	UserAgent    UserAgent `json:"userAgent"`
}

type SendMessage struct {
	CID         int64  `json:"cid"`
	DetectShare bool   `json:"detectShare"`
	Text        string `json:"text"`
}

type SendMessagePayload struct {
	ChatID  int         `json:"chatId"`
	Message SendMessage `json:"message"`
	Notify  bool        `json:"notify"`
	Type    string      `json:"type"`
}

type NotifyMessage struct {
	Sender   int64         `json:"sender"`
	ID       string        `json:"id"`
	Time     int           `json:"time"`
	Text     string        `json:"text"`
	Type     string        `json:"type"`
	CID      int64         `json:"cid"`
	Attaches []interface{} `json:"attaches"`
}

type NotifyMessagePayload struct {
	ChatID        int           `json:"chatId"`
	Message       NotifyMessage `json:"message"`
	TTL           bool          `json:"ttl"`
	Type          string        `json:"type"`
	PrevMessageID string        `json:"prevMessageId"`
}
