package i18n

import (
	"ttgoer/tiktok"

	"gopkg.in/telebot.v3"
)

func InvalidTikTokURL(c telebot.Context) string {
	switch c.Sender().LanguageCode {
	case "ua":
		return "Некоректне ТікТок посилання"

	default:
		return "Invalid TikTok URL"
	}
}

func UnknownError(c telebot.Context) string {
	switch c.Sender().LanguageCode {
	case "ua":
		return "Помилка :("

	default:
		return "Error :("
	}
}

func Error(error string, c telebot.Context) string {
	switch c.Sender().LanguageCode {
	default:
		return "Error: " + error
	}
}

func Welcome(c telebot.Context) string {
	switch c.Sender().LanguageCode {
	case "ua":
		return "Привіт! Відправляй мені коректні посилання на ТікТок, наприклад " + tiktok.ExampleTikTokURL

	default:
		return "Hello! Send me a valid TikTok URL, for example " + tiktok.ExampleTikTokURL
	}
}
