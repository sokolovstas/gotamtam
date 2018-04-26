package gotamtam

type Wrapper struct {
	Ver     int         `json:"ver"`
	Cmd     int         `json:"cmd"`
	Seq     int         `json:"seq"`
	OpCode  int         `json:"opcode"`
	Payload interface{} `json:"payload"`
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
	ChatsSync    int       `json:"chatsSync"`
	ConfigHash   string    `json:"configHash"`
	ContactSync  int       `json:"contactsSync"`
	PresenceSync int       `json:"presenceSync"`
	Token        string    `json:"token"`
	UserAgent    UserAgent `json:"userAgent"`
}

type Message struct {
	CID         int64  `json:"cid"`
	DetectShare bool   `json:"detectShare"`
	Text        string `json:"text"`
}

type MessagePayload struct {
	ChatID  int     `json:"chatId"`
	Message Message `json:"message"`
	Notify  bool    `json:"notify"`
	Type    string  `json:"type"`
}
