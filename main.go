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
		if update.Message != nil {			  
    
    			m := map[int]string{0:"🌖 Good night, ", 1:"🌅 Good morning, ",
                        		    2:"🌞 Good day, ", 3:"🌃 Good evening, "}

    			greeting = (m[hour/6])
	
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, greeting+update.Message.From.FirstName+welcomeMsg)
			bot.Send(msg)

		}
	}
}
