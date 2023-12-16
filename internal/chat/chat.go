package chat

import (
	"github.com/ASA11599/ws-chat/internal/set"
	"github.com/gofiber/websocket/v2"
)

type Chat struct {
	roomConnections map[string]set.Set[*websocket.Conn]
}

func NewChat() *Chat {
	return &Chat{
		roomConnections: make(map[string]set.Set[*websocket.Conn], 0),
	}
}

func (c *Chat) AddConnToRoom(room string, conn *websocket.Conn) {
	_, ok := c.roomConnections[room]
	if !ok {
		c.roomConnections[room] = set.NewHashSet[*websocket.Conn]()
	}
	c.roomConnections[room].Insert(conn)
}

func (c *Chat) DeleteConnFromRoom(room string, conn *websocket.Conn) {
	conns, ok := c.roomConnections[room]
	if ok { conns.Delete(conn) }
	if c.roomConnections[room].Size() == 0 {
		delete(c.roomConnections, room)
	}
}

func (c *Chat) RoomConns(room string) []*websocket.Conn {
	return c.roomConnections[room].Items()
}

func (c *Chat) Rooms() []ChatRoom {
	rooms := make([]ChatRoom, 0, len(c.roomConnections))
	for roomName, conns := range c.roomConnections {
		rooms = append(rooms, ChatRoom{ Name: roomName, Size: conns.Size() })
	}
	return rooms
}

func (c *Chat) Broadcast(t int, m []byte, room string) {
	for _, wsc := range c.RoomConns(room) {
		wsc.WriteMessage(t, m)
	}
}
