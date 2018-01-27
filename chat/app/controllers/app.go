package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	/*
		var user = models.User{
			Username: "Romi",
			Email:    "romi@iver.mx",
			Password: "hola",
		}
		revel.INFO.Printf("User %v", user.Username)
	*/
	return c.Render()
}

func (c App) Platzi(user, messageType string) revel.Result {
	revel.INFO.Printf("Dentro de Platzi")

	// Validación
	c.Validation.Required(user)
	c.Validation.Required(messageType)

	if c.Validation.HasErrors() {
		// Este es un mensaje de error Flash
		c.Flash.Error("Falta algún parámetro")
		// Para salir y redireccionar al Index
		return c.Redirect(App.Index)
	}

	switch messageType {
	case "refresh":
		return c.Redirect("/refresh?user=%s", user)
	case "polling":
		return c.Redirect("/polling/room?user=%s", user)
	case "websocket":
		return c.Redirect("/websocket/room?user=%s", user)
	}

	// Van a renderear los archivos estáticos de routes
	return c.Render()
	//return c.RenderError(errors.New("No está la vista"))
}
