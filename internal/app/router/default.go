package router

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *CommandRouter) Default(inputMessage *tgbotapi.Message) error {
	fmt.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote "+inputMessage.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	registeredCommands["default"] = (*CommandRouter).Default
}
