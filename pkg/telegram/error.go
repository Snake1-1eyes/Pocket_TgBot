package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url if invalid!")
	errUnauthorized = errors.New("user is not authorized!")
	errUnableToSave = errors.New("unable to save link=(")
)

//msg.Text = "Ты не авторизирован! Используй команду /start"
//msg.Text = "Это невалидная ссылка!"
//msg.Text = "Не удалось сохранить ссылку=( Попробуй еще раз позже"

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Default)

	switch err {
	case errInvalidURL:
		msg.Text = b.messages.InvalidURL
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = b.messages.Unauthorized
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = b.messages.UnableToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
