
# ResponseChatGPT

Description:
ResponseChatGPT is an innovative chatbot powered by Golang and OpenAI's advanced language model, designed to provide intelligent and interactive conversation capabilities. It leverages the cutting-edge capabilities of Golang and OpenAI's language model to generate contextually appropriate responses in natural language.

Features:
- Seamless integration with Golang for high-performance and efficient execution
- Utilizes OpenAI's powerful language model to generate human-like responses
- Supports multi-turn conversations for interactive and dynamic chat experiences
- Easy-to-use API for integrating the chatbot into your applications or systems

## Authors

- [@ReqqQ](https://www.github.com/ReqqQ)


## API Reference

#### Get info about weather

```http
  POST /api/v1/response/message/generateTooltip
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `message` | `string` | **Required**. Message to find context |

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`CHAT_GPT_KEY`


## Tech Stack

**Client API:** Fast HTTP

**Lang:** GoLang
