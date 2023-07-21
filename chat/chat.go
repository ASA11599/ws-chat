package chat

import (
	"github.com/ASA11599/ws-chat/set"
	"github.com/gofiber/websocket/v2"
)

type Chat map[string]set.Set[*websocket.Conn]

func NewChat() Chat {
	return make(map[string]set.Set[*websocket.Conn], 0)
}

func (c *Chat) AddConnToRoom(room string, conn *websocket.Conn) {
	_, ok := (*c)[room]
	if !ok {
		(*c)[room] = set.NewHashSet[*websocket.Conn]()
	}
	(*c)[room].Insert(conn)
}

func (c *Chat) DeleteConnFromRoom(room string, conn *websocket.Conn) {
	conns, ok := (*c)[room]
	if ok {
		conns.Delete(conn)
	}
}

func (c *Chat) RoomConns(room string) []*websocket.Conn {
	return (*c)[room].Items()
}
