package DomainModelChatGPT

import "encoding/json"

type DomainModelChatGPTBuilder struct {
}

func (builder DomainModelChatGPTBuilder) BuildMessagePrompt(vo MessageVO) json.RawMessage {
	content := "Jak można odpowiedzieć na to : '" + vo.GetContent() + "' .Odpowiedź ostylowana w HTML"
	encoded, _ := json.Marshal(content)

	return encoded
}
