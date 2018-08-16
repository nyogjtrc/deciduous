package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Websocket(router *gin.Engine) {
	router.GET("/ws", func(c *gin.Context) {
	})
}

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandlero(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err.Error())
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read messag error:", err.Error())
			break
		}
		fmt.Println(message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write message error:", err.Error())
			break
		}
	}
}
