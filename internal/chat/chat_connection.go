package chat

import "github.com/gofiber/websocket/v2"

type ChatConnection interface {
	Read() (ChatMessage, error, bool)
	Write(msg ChatMessage) error
	Close() error
	Client() string
}

type WSChatConnection struct {
	wsConn *websocket.Conn
}

func NewWSChatConnection(wsc *websocket.Conn) WSChatConnection {
	return WSChatConnection{
		wsConn: wsc,
	}
}

func (wscc WSChatConnection) Client() string {
	return wscc.wsConn.RemoteAddr().String()
}

func (wscc WSChatConnection) Read() (ChatMessage, error, bool) {
	typ, msg, err := wscc.wsConn.ReadMessage()
	return WSChatMessage{ typ: typ, content: msg }, err, websocket.IsUnexpectedCloseError(
		err,
		websocket.CloseMessage,
		websocket.CloseGoingAway,
		websocket.CloseNormalClosure,
		websocket.CloseNoStatusReceived,
	)
}

func (wscc WSChatConnection) Write(msg ChatMessage) error {
	wscm, ok := msg.(WSChatMessage)
	if ok {
		return wscc.WriteWS(wscm)
	} else {
		return wscc.wsConn.WriteMessage(websocket.BinaryMessage, msg.Content())
	}
}

func (wscc WSChatConnection) WriteWS(msg WSChatMessage) error {
	return wscc.wsConn.WriteMessage(msg.typ, msg.Content())
}

func (wscc WSChatConnection) Close() error {
	return wscc.wsConn.Close()
}
