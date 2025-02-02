package controllers

import (
	"v/pkg/config"
	"v/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/golang-jwt/jwt/v5"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/webhook"
)

func HandleWebhook(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id,ok := claims["user_id"].(string)

	admin := c.Query("admin") == "true"
	if !ok {
		return fiber.ErrUnauthorized
	}

	authProvider := auth.NewSimpleKeyProvider(
	    config.Conf.Livekit.ApiKey, config.Conf.Livekit.Secret,
	  )
	req,err :=  adaptor.ConvertRequest(c, true) 
	if err != nil {
		return fiber.ErrInternalServerError
	}

	  // event is a livekit.WebhookEvent{} object
	event, err := webhook.ReceiveWebhookEvent(req, authProvider)
	if err != nil {
		println(err)
		return err
	}
	if event.Event == webhook.EventParticipantJoined {

		if event.Participant.Identity == config.BotIdentity {
			return nil
		}
		rm := models.NewRoomService(config.App)

		rm.JoinRoom(user_id, config.Conf.Livekit.Host, admin, event.Room)

	}

	return nil
}
