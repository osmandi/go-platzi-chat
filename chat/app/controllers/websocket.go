package controllers

import (
	"platzi/chat/app/lib/chat"

	"github.com/revel/revel"
	"golang.org/x/net/websocket"
)

// Parte del controlador de revel
type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) Room(user string) revel.Result {
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws *websocket.Conn) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()
	chat.Join(user)
	defer chat.Leave(user)

	// Send down the archive
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				return nil
			}
		case msg, ok := <-newMessages:
			if ok == false {
				return nil
			}

			chat.Say(user, msg)
		}
	}
	return nil
}
