package AppMessage

import (
	"github.com/gofiber/fiber/v2"

	DomainModelChatGPT "ResponseChatGPT/DomainModel/ChatGPT"
)

type MessageFactory struct {
}

type MessageTooltipCommand struct {
	Message string `json:"message" form:"message"`
}

type MessageTooltipDTO struct {
	Message string
}

func (factory MessageFactory) GetMessageTooltipCommand(c *fiber.Ctx) MessageTooltipCommand {
	command := new(MessageTooltipCommand)
	err := c.BodyParser(command)
	if err != nil {
		return MessageTooltipCommand{}
	}

	return *command
}

func (factory MessageFactory) GetMessageTooltipDTO(message DomainModelChatGPT.MessageVO) MessageTooltipDTO {
	return MessageTooltipDTO{
		Message: message.GetContent(),
	}
}

func (factory MessageFactory) GetMessageVO(command MessageTooltipCommand) DomainModelChatGPT.MessageVO {
	return DomainModelChatGPT.MessageVO{Content: command.Message}
}
