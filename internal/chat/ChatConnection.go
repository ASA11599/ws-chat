package chat

import "github.com/gofiber/websocket/v2"

type ChatConnection interface {
	Write(msg []byte) error
}

type WSChatConnection struct {
	wsConn *websocket.Conn
}

func NewWSChatConnection(wsc *websocket.Conn) *WSChatConnection {
	return &WSChatConnection{
		wsConn: wsc,
	}
}

func (wscc *WSChatConnection) Write(msg []byte) error {
	return wscc.wsConn.WriteMessage(websocket.TextMessage, msg)
}
