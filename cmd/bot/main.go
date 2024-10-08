package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/npavlov/telegram-bot-go/internal/app/router"
	"github.com/npavlov/telegram-bot-go/internal/config"
	"github.com/npavlov/telegram-bot-go/internal/service/product"
	"log"
	"os"
)

func main() {
	cfg := config.NewConfigBuilder().FromEnv().Build()

	fmt.Println(os.Getenv("TOKEN"))

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	commander := router.NewRouter(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
