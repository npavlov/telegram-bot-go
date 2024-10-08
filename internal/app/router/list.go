package router

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *CommandRouter) List(inputMessage *tgbotapi.Message) error {
	outputMessage := "Here all the products: \n\n"

	products := c.service.List()

	for _, p := range products {
		outputMessage += p.Title
		outputMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	serializedData, _ := json.Marshal(CommandData{Offset: 10})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData))))

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	registeredCommands["list"] = (*CommandRouter).List
}
