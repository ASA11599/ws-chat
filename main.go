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
	ch := chat.NewChat(log.Default())
	app.Get("/rooms", func(c *fiber.Ctx) error {
		return c.JSON(ch.Rooms())
	})
	app.Get("/:room/ws", websocket.New(func(c *websocket.Conn) {
		room := c.Params("room")
		ch.HandleConnection(room, chat.NewWSChatConnection(c))
	}))
	go app.Listen(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	<-interruptChannel()
}

