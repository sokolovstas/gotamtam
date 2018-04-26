# gotamtam
mail.ru TamTam Golang API

```
b, err := gotamtam.ComposeMessage(seq, opCode, payload)
if err != nil {
  log.Println("Compose error:", err)
  return
}
fmt.Println(string(b))
err = c.WriteMessage(websocket.TextMessage, b)
if err != nil {
  log.Println("Write error:", err)
  return
}
seq++
```
