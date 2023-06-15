package AppMessage

import DomainModelChatGPT "ResponseChatGPT/DomainModel/ChatGPT"

type MessageService struct {
	ChatGPTService DomainModelChatGPT.ChatGPTService
	Factory        MessageFactory
}

func (service MessageService) GetMessageTooltip(command MessageTooltipCommand) MessageTooltipDTO {
	vo := service.Factory.GetMessageVO(command)

	return service.Factory.GetMessageTooltipDTO(service.ChatGPTService.GetMessageTooltip(vo))
}
