package websocket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
}

type AppMsgReq struct {
	MsgType string `json:"msgType"`
	MsgData string `json:"msgData"`
}

type AppMsgResp struct {
	MsgType string   `json:"msgType"`
	MsgData []string `json:"msgData"`
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	var funds map[string]bool = make(map[string]bool)

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// 接收消息
		message := Message{Type: messageType, Body: string(p)}
		fmt.Printf("Message Received: %+v\n", message)
		var appMsg AppMsgReq
		if err := json.Unmarshal(p, &appMsg); err != nil {
			log.Println(err)
		}

		// 业务逻辑处理
		if appMsg.MsgType == "Add" {
			funds[appMsg.MsgData] = true
		} else if appMsg.MsgType == "Del" {
			delete(funds, appMsg.MsgData)
		}

		// ------------------
		// 暂时留空，固定响应
		currentDate := time.Now().Format(time.RFC3339)
		startDate := currentDate
		interval := -30
		endDate := time.Now().AddDate(0, 0, interval).Format(time.RFC3339)
		url := "http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz&code=000001&sdate=" + startDate + "&edate=" + endDate + "&per=50&page=1"
		fmt.Println(url)

		// 发送GET请求
		resp, err := http.Get(url)
		if err != nil {
			// 处理错误
			panic(err)
		}

		// 确保在函数退出时关闭响应体
		defer resp.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// 处理错误
			panic(err)
		}

		// 打印网页内容
		fmt.Println(string(len(body)))

		// 返回响应
		// c.Conn.WriteJSON(Message{Type: 1, Body: string(body)})

		// ------------------
		var appMsgResp AppMsgResp
		appMsgResp.MsgType = "general"
		for fund := range funds {
			appMsgResp.MsgData = append(appMsgResp.MsgData, fund)
		}
		jsonStr, err := json.Marshal(appMsgResp)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Message Send: %+v\n", appMsgResp)
		c.Conn.WriteJSON(Message{Type: 1, Body: string(jsonStr)})
	}
}
