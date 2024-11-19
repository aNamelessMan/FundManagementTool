package main

import (
	"fmt"
	"net/http"

	"github.com/aNamelessMan/FundManagementTool/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	// 直接postman/ws连调试不了，不知道为啥
	// 但是http://www.jsons.cn/websocket/ 和 前端都能正常连上
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Fund Tool App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
