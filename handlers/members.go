package handlers

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/DeveloperDurp/durpot/model"
)

var (
	ChannelID = model.ChannelID
)

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	message := fmt.Sprintf("Welcome <@%s> to our server!", m.Member.User.ID)
	_, err := s.ChannelMessageSend(ChannelID, message)
	if err != nil {
		log.Printf("Error sending welcome message: %v\n", err)
	}
}

func GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	message := fmt.Sprintf("Goodbye %s", m.Member.User.Username)
	_, err := s.ChannelMessageSend(ChannelID, message)
	if err != nil {
		log.Printf("Error sending goodbye message: %v\n", err)
	}
}
