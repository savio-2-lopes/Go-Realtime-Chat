package backend

import "home/savio/go_simple_realtime_chat/backend/utils"// modifique este caminho com o diretorio da pasta

type Message struct {
	ID int64 `json:"id"`
	Body string `json:"body"`
	Sender string `json:"sender"`
}

func NewMessage(body string, sender string) *Message  {
	return &Message{
		ID: utils.GetRandomI64(),
		Body: body,
		Sender: sender,
	}
}
