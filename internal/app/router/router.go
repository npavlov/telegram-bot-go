package router

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/npavlov/telegram-bot-go/internal/service/product"
	"log"
)

type CommandData struct {
	Offset int `json:"offset"`
}

var registeredCommands = map[string]func(c *CommandRouter, message *tgbotapi.Message) error{}

type CommandRouter struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewRouter(bot *tgbotapi.BotAPI, productService *product.Service) *CommandRouter {
	return &CommandRouter{
		bot:     bot,
		service: productService,
	}
}
func (c *CommandRouter) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovered from panic %v", err)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		_ = json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)

		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed %+v\n", parsedData))
		_, _ = c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	command, ok := registeredCommands[update.Message.Command()]

	if ok {
		err := command(c, update.Message)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := c.Default(update.Message)
		if err != nil {
			log.Println(err)
		}
	}
}
