package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/use-go/websocket-streamserver/logger"
	"github.com/use-go/websocket-streamserver/webSocketService"
	"github.com/use-go/websocket-streamserver/wssAPI"
)

func main() {
	logger.SetFlags(logger.LOG_SHORT_FILE)
	cli := &websocket.Dialer{}
	req := http.Header{}

	conn, _, err := cli.Dial("ws://127.0.0.1:8080/ws/live", req)
	if err != nil {
		logger.LOGE(err.Error())
		return
	}
	Play(conn)
	defer conn.Close()

}

// SetTest test
func SetTest() {
	si := wssAPI.NewSet()
	si.Add("111")
	si.Add(2)
	fmt.Println(si.Has(3))
	fmt.Println(si.Has(2))
	fmt.Println(si.Has("111"))
}

type stPlay struct {
	Name  string `json:"name"`
	Start int    `json:"start"`
	Len   int    `json:"len"`
	Reset int    `json:"reset"`
	Req   int    `json:"req"`
}

// Play video
func Play(conn *websocket.Conn) {

	//play
	stPlay := &stPlay{}

	stPlay.Name = "hks"

	dataSend, _ := json.Marshal(stPlay)
	err := webSocketService.SendWsControl(conn, webSocketService.WSCPlay, dataSend)
	if err != nil {
		logger.LOGE(err.Error())
		return
	}
	err = readResult(conn)
	if err != nil {
		logger.LOGE(err.Error())
		return
	}
	err = readResult(conn)
	if err != nil {
		logger.LOGE(err.Error())
		return
	}
	for {
		readResult(conn)
	}
}

func readResult(conn *websocket.Conn) (err error) {
	msgType, data, err := conn.ReadMessage()
	if err != nil {
		logger.LOGE(err.Error())
		return
	}
	if msgType == websocket.BinaryMessage {
		logger.LOGT(data)
		pktType := data[0]
		switch pktType {
		case webSocketService.WSPktAudio:
			logger.LOGT("audio")
		case webSocketService.WSPktVideo:
			logger.LOGT("video")
		case webSocketService.WSPktControl:
			logger.LOGT(data)
		}
	}
	return
}
