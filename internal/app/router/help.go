package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *CommandRouter) Help(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Usage: /help\n"+
		"list - list products")

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	registeredCommands["help"] = (*CommandRouter).Help
}
