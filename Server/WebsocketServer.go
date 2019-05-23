package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func ping(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		log.Println(string(message))
		if string(message) == "ping" {
			message = []byte("pong")
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
		
	}

}
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/ping",ping)
	if err := router.Run(":2303");err != nil{
		log.Println(err.Error())
	}
}