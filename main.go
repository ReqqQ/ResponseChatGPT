package main

import (
	"ResponseChatGPT/Config"
	"ResponseChatGPT/UI"
)

func main() {
	app := Config.Init()
	controller := new(UI.Controllers)
	controller.InitGetRoutes(app)
	Config.Start(app)
}
