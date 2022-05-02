package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var greeting string
var welcomeMsg string

func main() {
	hour := time.Now().Hour()

	bot, err := tgbotapi.NewBotAPI("TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		b, err := ioutil.ReadFile("users.txt")
		str := string(b)
		if err != nil {
			log.Fatal(err)
		}
		user_name := []byte(strconv.FormatInt(update.Message.From.ID, 10) + "\n")
		if strings.Contains(str, string(user_name)) {
			println("USER ALREADY REGISTERED")
			welcomeMsg = " 🔑 Welcome back!"

		} else {
			f, err := os.OpenFile("users.txt", os.O_RDWR|os.O_APPEND, 0660)
			_, err2 := f.Write(user_name)
			println("REGISTERED NEW USER " + string(user_name))
			welcomeMsg = " 🦆 Now you are a new member of the club!"
			if err2 != nil {
				log.Fatal(err2)
			}
			if err != nil {
				log.Fatal(err)
			}
		}
		if update.Message != nil { // If we got a message
			//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch {
			case 0 <= hour && hour < 6:
				greeting = "🌖 Good night, "
			case 6 <= hour && hour < 12:
				greeting = "🌅 Good morning, "
			case 12 <= hour && hour < 17:
				greeting = "🌞 Good day, "
			case 17 <= hour && hour <= 23:
				greeting = "🌃 Good evening, "
			default:
				greeting = strconv.Itoa(hour)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, greeting+update.Message.From.FirstName+welcomeMsg)
			bot.Send(msg)

		}
	}
}
