package controllers

import (
	"platzi/chat/app/lib/chat"
	"platzi/chat/app/routes"

	"github.com/revel/revel"
)

type Refresh struct {
	*revel.Controller
}

func (c *Refresh) Index(user string) revel.Result {
	chat.Join(user)
	revel.INFO.Printf("Dentro de Index - Room", user)

	// Para que envíe la ruta a Room
	return c.Redirect(routes.Refresh.Room(user))
	//return c.Render()

}

func (c *Refresh) Room(user string) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()
	events := subscription.Archive
	for i, _ := range events {
		if events[i].User == user {
			events[i].User = "tu"
		}
	}
	return c.Render(user, events)

}

func (c *Refresh) Say(user, message string) revel.Result {
	revel.INFO.Printf("Usuario: %v message: %v", user, message)
	chat.Say(user, message)
	return c.Redirect(routes.Refresh.Room(user))
}

func (c *Refresh) Leave(user string) revel.Result {
	chat.Leave(user)
	return c.Redirect(App.Index)
}
