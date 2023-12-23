package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ASA11599/ws-chat/internal/chat"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func interruptChannel() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	return c
}

func main() {
	app := fiber.New()
	defer app.Shutdown()
	ch := chat.NewChat()
	logger := log.Default()
	app.Get("/rooms", func(c *fiber.Ctx) error {
		return c.JSON(ch.Rooms())
	})
	app.Get("/:room/ws", websocket.New(func(c *websocket.Conn) {
		room := c.Params("room")
		logger.Println("Client", c.RemoteAddr(), "connected to room", room)
		defer func() {
			logger.Println("Client", c.RemoteAddr(), "disconnected from room", room)
			ch.DeleteConnFromRoom(room, chat.NewWSChatConnection(c))
			c.Close()
		}()
		ch.AddConnToRoom(room, chat.NewWSChatConnection(c))
		for {
			typ, msg, err := c.ReadMessage()
			if (err != nil) {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
					logger.Println("Error reading message:", err)
				}
				break
			}
			if (typ == websocket.BinaryMessage) || (typ == websocket.TextMessage) {
				logger.Println("Client", c.RemoteAddr(), "sent", string(msg), "to room", room)
				ch.Broadcast(msg, room)
			}
		}
	}))
	go app.Listen(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	<-interruptChannel()
}

