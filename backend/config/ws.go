package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[string]*websocket.Conn)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Wshandler(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	userID := c.Query("id")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	// _, username := model.CheckUserExist(userID)
	// if username == "" {
	// Closews("Authentication failed - invalid username", conn)
	// return
	// }

	NewClient(userID, conn)
	ReceiveMessage(conn, userID)

}

func NewClient(userId string, conn *websocket.Conn) {
	clients[userId] = conn
	clients[userId].WriteMessage(websocket.TextMessage, []byte("ok"))
}

type Message struct {
	Action string `json:"action"`
}

func ReceiveMessage(conn *websocket.Conn, userid string) {

	for {
		_, msg, errCon := conn.ReadMessage()

		if errCon != nil {
			log.Println("Read Error:", errCon)
			break
		}
		var r Message
		if err := json.Unmarshal(msg, &r); err != nil {
			log.Println("Error: " + err.Error())
			MsgFailed(conn)
			continue
		}
		fmt.Println("Message struct: ", r)
		broadcast <- r

	}
}

func Send() {
	for {
		msg := <-broadcast
		fmt.Println("message SEND:", msg)
		for i, conn := range clients {
			fmt.Println("connection: ",i)
			jsonData, _ := json.Marshal(msg)
			conn.WriteMessage(websocket.TextMessage, jsonData)
		}
	}
}

func MsgFailed(conn *websocket.Conn) {
	msg := `{"message":"Failed to send message"}`
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		fmt.Println(err)
	}
}
