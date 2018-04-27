# GoTamTam
TamTam Golang API

```
func main() {
	tamtam, err := gotamtam.New(TOKEN, "0.1.0", "TestBOT")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	tamtam.Serve(NewGiphyBot())
}

type GiphyBot struct {
	giphy *libgiphy.Giphy
}

func NewGiphyBot() *GiphyBot {
	b := &GiphyBot{}
	b.giphy = libgiphy.NewGiphy("")
	return b
}

func (b *GiphyBot) Response(client *gotamtam.Client, message *gotamtam.Message) {
	switch message.OpCode {
	case gotamtam.NOTIF_MESSAGE:
		m := message.Payload.(gotamtam.NotifyMessagePayload)

		dataTranslate, err := b.giphy.GetTranslate(m.Message.Text, "", "", false)
		if err == nil && dataTranslate != nil {
			client.SendMessage(m.ChatID, m.Type, dataTranslate.Data.Images.Downsized.Url)
		}
	}
}

```
