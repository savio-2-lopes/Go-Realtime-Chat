package backend	

import (
	"home/savio/go_simple_realtime_chat/backend/utils" // modifique este caminho com o diretorio da pasta
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"log"
	"fmt"
)

type Chat struct {
	users map[string]*User
	messages chan *Message
	join chan *User
	leave chan *User
}

var upgrader = websocket.Upgrader {
	ReadBufferSize: 512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		log.Printf("%s %s%s %v\n", r.Method, r.Host, r.RequestURI, r.Proto)
		return r.Method == http.MethodGet
	},
}

func (c *Chat) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("Erro na conexão:", err.Error())
	}

	keys := r.URL.Query()
	username := keys.Get("username")
	if strings.TrimSpace(username) == "" {
		username = fmt.Sprintf("anon-%d", utils.GetRandomI64())
	}

	user := &User {
		Username: username,
		Conn: conn,
		Global: c,
	}

	c.join <- user

	user.Read()
}

func (c *Chat) Run() {
	for {
		select {
		case user := <-c.join:
			c.add(user)
		case message := <-c.messages:
			c.broadcast(message)
		case user := <-c.leave:
			c.disconnect(user)
		}
	}
}

func (c *Chat) add(user *User) {
	if _, ok := c.users[user.Username]; !ok {
		c.users[user.Username] = user
		body := fmt.Sprintf(" %s entrou no PetChat", user.Username)
		c.broadcast(NewMessage(body, "PetChat"))
	}
}

func (c *Chat) broadcast(message *Message) {
	log.Printf("Mensagem transmitida: %v\n", message)
	for _, user := range c.users {
		user.Write(message)
	}
}

func (c *Chat) disconnect(user *User) {
	if _, ok := c.users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.users, user.Username)
		body := fmt.Sprintf(" %s saiu do PetChat", user.Username)
		c.broadcast(NewMessage(body, "PetChat"))
	}
}

func Start(port string) {
	log.Printf("🐾 PetChat aberto em http://localhost%s\n", port)
	c := &Chat{
		users: make(map[string]*User),
		messages: make(chan *Message),
		join: make(chan *User),
		leave: make(chan *User),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem vindo ao 🐾 PetChat!"))
	})

	http.HandleFunc("/chat", c.Handler)
	go c.Run()
	log.Fatal(http.ListenAndServe(port, nil))
}