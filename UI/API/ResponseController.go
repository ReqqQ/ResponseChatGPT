package API

import (
	"github.com/gofiber/fiber/v2"

	AppMessage "ResponseChatGPT/App/Message"
)

type ResponseController struct {
	MessageService AppMessage.MessageService
	MessageFactory AppMessage.MessageFactory
}

func (wr ResponseController) GetMessageTooltip(c *fiber.Ctx) AppMessage.MessageTooltipDTO {
	command := wr.MessageFactory.GetMessageTooltipCommand(c)

	return wr.MessageService.GetMessageTooltip(command)
}
