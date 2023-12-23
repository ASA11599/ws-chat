package chat

import (
	"github.com/ASA11599/ws-chat/internal/set"
)

type Chat struct {
	roomConnections map[string]set.Set[ChatConnection]
}

func NewChat() *Chat {
	return &Chat{
		roomConnections: make(map[string]set.Set[ChatConnection], 0),
	}
}

func (c *Chat) AddConnToRoom(room string, conn ChatConnection) {
	_, ok := c.roomConnections[room]
	if !ok {
		c.roomConnections[room] = set.NewHashSet[ChatConnection]()
	}
	c.roomConnections[room].Insert(conn)
}

func (c *Chat) DeleteConnFromRoom(room string, conn ChatConnection) {
	conns, ok := c.roomConnections[room]
	if ok { conns.Delete(conn) }
	if c.roomConnections[room].Size() == 0 {
		delete(c.roomConnections, room)
	}
}

func (c *Chat) RoomConns(room string) []ChatConnection {
	return c.roomConnections[room].Items()
}

func (c *Chat) Rooms() []ChatRoom {
	rooms := make([]ChatRoom, 0, len(c.roomConnections))
	for roomName, conns := range c.roomConnections {
		rooms = append(rooms, ChatRoom{ Name: roomName, Size: conns.Size() })
	}
	return rooms
}

func (c *Chat) Broadcast(m []byte, room string) {
	for _, conn := range c.RoomConns(room) {
		conn.Write(m)
	}
}
