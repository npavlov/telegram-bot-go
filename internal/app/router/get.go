package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *CommandRouter) Get(message *tgbotapi.Message) error {
	inputArgs := message.CommandArguments()

	arg, err := strconv.Atoi(inputArgs)

	if err != nil {
		log.Println("wrong args", err)
		return err
	}

	product, err := c.service.Get(arg)

	if err != nil {
		log.Printf("product not found with id %d: %v", arg, err)
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		product.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	registeredCommands["get"] = (*CommandRouter).Get
}
