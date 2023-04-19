package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
	"github.com/vyroai/VyroAI/commons/errors"
)

type messageRequest struct {
	Content string `json:"content"`
}

func (s *WebServiceHttpServer) createMessage(c *fiber.Ctx) error {
	var messageRequestPayload messageRequest

	if err := c.BodyParser(&messageRequestPayload); err != nil {
		response.ErrorJson(c, 400, err.Error())
		return nil
	}

	if messageRequestPayload.Content == "" {
		response.ErrorJson(c, 400, "content is required")
		return nil
	}
	chatBotID, err := c.ParamsInt("chatbotID")
	message, err := s.chatService.CreateMessage(c.Context(), messageRequestPayload.Content, int64(chatBotID))
	if err != nil {
		switch errors.GetType(err) {
		default:
			response.ServerError(c)
			return nil
		}
	}

	response.SuccessDataJson(c, 201, message)
	return nil
}

func (s *WebServiceHttpServer) allMessages(c *fiber.Ctx) error {
	chatBotID, err := c.ParamsInt("chatbotID")

	limit := c.QueryInt("limit")
	if limit == 0 {
		limit = 50
	}
	offset := c.QueryInt("offset")
	if offset == 0 {
		offset = 1
	}

	messages, err := s.chatService.GetChatMessagesByChatID(
		c.Context(),
		int64(chatBotID),
		int32(limit),
		int32(offset))

	if err != nil {
		switch errors.GetType(err) {
		default:
			response.ServerError(c)
			return nil
		}
	}

	response.SuccessDataJson(c, 200, messages)
	return nil
}
