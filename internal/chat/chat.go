package chat

import (
	"log"

	"github.com/ASA11599/ws-chat/internal/set"
)

type Chat struct {
	roomConnections map[string]set.Set[ChatConnection]
	logger *log.Logger
}

func NewChat(logger *log.Logger) *Chat {
	return &Chat{
		roomConnections: make(map[string]set.Set[ChatConnection], 0),
		logger: logger,
	}
}

func (wsc *Chat) addConnToRoom(room string, conn ChatConnection) {
	_, ok := wsc.roomConnections[room]
	if !ok {
		wsc.roomConnections[room] = set.NewHashSet[ChatConnection]()
	}
	wsc.roomConnections[room].Insert(conn)
}

func (wsc *Chat) deleteConnFromRoom(room string, conn ChatConnection) {
	conns, ok := wsc.roomConnections[room]
	if ok {
		conns.Delete(conn)
		if conns.Size() == 0 {
			delete(wsc.roomConnections, room)
		}
	}
}

func (wsc *Chat) getRoomConnections(room string) []ChatConnection {
	conns, ok := wsc.roomConnections[room]
	if ok {
		return conns.Items()
	} else {
		return nil
	}
}

func (wsc *Chat) Rooms() []ChatRoom {
	rooms := make([]ChatRoom, 0, len(wsc.roomConnections))
	for roomName, conns := range wsc.roomConnections {
		rooms = append(rooms, ChatRoom{ Name: roomName, Size: conns.Size() })
	}
	return rooms
}

func (wsc *Chat) broadcast(m ChatMessage, room string) {
	for _, conn := range wsc.getRoomConnections(room) {
		conn.Write(m)
	}
}

func (wsc *Chat) HandleConnection(room string, conn ChatConnection) {
	wsc.logger.Printf("Client %s connected to room %s", conn.Client(), room)
	defer wsc.logger.Printf("Client %s disconnected to room %s", conn.Client(), room)
	defer conn.Close()
	wsc.addConnToRoom(room, conn)
	defer wsc.deleteConnFromRoom(room, conn)
	for {
		msg, err, open := conn.Read()
		if (err != nil) {
			wsc.logger.Printf("Error reading from connection: %s", err.Error())
			if !open { break }
		} else {
			wsc.logger.Printf("Client %s sent %s to room %s", conn.Client(), string(msg.Content()), room)
			wsc.broadcast(msg, room)
		}
	}
}
