package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	openai "github.com/sashabaranov/go-openai"
	"gitlab.com/DeveloperDurp/durpot/handlers"
	"gitlab.com/DeveloperDurp/durpot/model"
)

var (
	Token     = model.Token
	BotPrefix = model.BotPrefix
	ChannelID = model.ChannelID
	BotId     string
	ApiKey    = model.ApiKey
)

func Start() {
	goBot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)
	goBot.AddHandler(handlers.GuildMemberAdd)
	goBot.AddHandler(handlers.GuildMemberRemove)
	goBot.AddHandler(handleTag)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	//baseurl := "https://kong.durp.info/"
	switch m.Content {
	case BotPrefix + "ping":
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Failed to send Message")
		}
	case BotPrefix + "unraid":
		handlers.GetUnraidUsage(s, m)

	}

}

func main() {

	Start()

	<-make(chan struct{})
}

func handleTag(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	client := openai.NewClient(ApiKey)

	// Check if bot is mentioned in message
	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			// Remove mention from message content
			content := strings.Replace(m.ContentWithMentionsReplaced(), "<@"+s.State.User.ID+">", "", -1)
			content = strings.Replace(content, "<@!"+s.State.User.ID+">", "", -1)
			content = strings.TrimSpace(content)

			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: content,
						},
					},
				},
			)

			if err != nil {
				fmt.Printf("ChatCompletion error: %v\n", err)
				return
			}

			fmt.Println(resp.Choices[0].Message.Content)

			// Send generated response back to Discord
			_, err = s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
