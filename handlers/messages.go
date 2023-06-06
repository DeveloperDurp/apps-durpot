package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/DeveloperDurp/durpot/model"
)

var (
	BotPrefix = model.BotPrefix
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case BotPrefix + "ping":
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Failed to send Message")
		}
	case BotPrefix + "unraid":
		GetUnraidUsage(s, m)

	}

}
