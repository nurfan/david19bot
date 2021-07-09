package model

import "os"

// Define a few constants and variable to handle different commands
const PunchCommand string = "/punch"

var LenPunchCommand int = len(PunchCommand)

const StartCommand string = "/start"

var LenStartCommand int = len(StartCommand)

const BotTag string = "@DavidAutoBot"

var LenBotTag int = len(BotTag)

// Pass token and sensible APIs through environment variables
var (
	telegramApiBaseUrl     string = os.Getenv("TELEGRAM_URL")
	telegramApiSendMessage string = "/sendMessage"
	telegramTokenEnv       string = os.Getenv("TELEGRAM_TOKEN")
	TelegramApiSend        string = telegramApiBaseUrl + telegramTokenEnv + telegramApiSendMessage
)

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object that can be found in an update.
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}
