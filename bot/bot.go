package bot

import (
	"bytes"

	"ttgoer/cfg"
	"ttgoer/i18n"
	"ttgoer/log"
	"ttgoer/tiktok"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

var bot *telebot.Bot

func Start() {
	log.S().Info("initializing bot")

	poller := telebot.LongPoller{Timeout: cfg.Get().Bot.PollerTimeout}
	settings := telebot.Settings{Token: cfg.Get().Bot.Token, Poller: &poller}

	var err error
	if bot, err = telebot.NewBot(settings); err != nil {
		log.S().Fatalf("failed to create a bot instance: %v", err)
	}

	log.S().Infof("Bot: %+v", bot.Me)

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		chat := c.Chat()

		if !cfg.Get().IsUserAllowed(c.Sender().ID) {
			return nil
		}

		if c.Text() == "/start" {
			return c.Reply(i18n.Welcome(c), &telebot.SendOptions{
				DisableWebPagePreview: true,
			})
		}

		if chat.Type == telebot.ChatPrivate {
			go handleDirectDownloadRequest(c)
			return nil
		}

		if chat.Type == telebot.ChatGroup || chat.Type == telebot.ChatSuperGroup {
			// TODO
			return nil
		}

		return nil
	})

	log.S().Info("started polling for the bot")
	go bot.Start()
}

func Stop() {
	log.S().Info("stopping the bot")
	bot.Stop()
}

func handleDirectDownloadRequest(c telebot.Context) {
	url := c.Text()

	if !tiktok.IsValidTikTokURL(url) {
		_ = c.Reply(i18n.InvalidTikTokURL(c))
		return
	}

	tt, err := tiktok.Fetch(url, c.Sender().Username)

	if err != nil || tt == nil {
		_ = c.Reply(i18n.Error(err.Error(), c))
		log.S().
			With(zap.String("username", c.Sender().Username)).
			With(zap.String("tikTokURL", url)).
			Errorf("request failed: %v", err)
		return
	}

	videoFile := &telebot.Video{
		File: telebot.FromReader(bytes.NewReader(tt.FileBytes)),
	}

	if err = c.Reply(videoFile); err != nil {
		log.S().
			With(zap.String("username", c.Sender().Username)).
			With(zap.String("tikTokURL", url)).
			Errorf("failed to send video: %v", err)
	} else {
		log.S().
			With(zap.String("username", c.Sender().Username)).
			With(zap.String("tikTokURL", url)).
			Info("sent video")
	}

}
