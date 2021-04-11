package backend 

import (
	"github.com/gorilla/websocket"
	"encoding/json" 
	"log"
)

type User struct {
	Username string
	Conn *websocket.Conn
	Global *Chat
}

func (u *User) Read() {
	for {
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Println("Error ao reconhecer a mensagem:", err.Error())

			break
		} else {
			u.Global.messages <- NewMessage(string(message), u.Username)
		}
	}
	u.Global.leave <- u
}

func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("Error ao escrever a mensagem:", err.Error())
	}
}
