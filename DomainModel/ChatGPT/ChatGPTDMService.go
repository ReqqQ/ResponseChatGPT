package DomainModelChatGPT

import (
	InfrastructureChatGPTPackage "ResponseChatGPT/Infrastructure/ChatGPTPackage"
)

type ChatGPTService struct {
	ChatGPTPackage InfrastructureChatGPTPackage.InfrastructureChatGPTPackage
	Factory        DomainModelChatGPTFactory
	Builder        DomainModelChatGPTBuilder
}

func (cg ChatGPTService) GetMessageTooltip(vo MessageVO) MessageVO {
	messageDTO := cg.ChatGPTPackage.GetChatGPTWheather(cg.Builder.BuildMessagePrompt(vo))

	return cg.Factory.GetMessageVO(messageDTO.Content)
}
