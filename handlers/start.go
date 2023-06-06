package handlers

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/DeveloperDurp/durpot/model"
)

func Start() error {
	goBot, err := discordgo.New("Bot " + model.Token)

	if err != nil {
		return (err)
	}

	goBot.AddHandler(MessageHandler)
	goBot.AddHandler(GuildMemberAdd)
	goBot.AddHandler(GuildMemberRemove)
	goBot.AddHandler(HandleTag)

	err = goBot.Open()

	if err != nil {
		return (err)
	}

	return (err)
}
