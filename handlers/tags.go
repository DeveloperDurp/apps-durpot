package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
	"gitlab.com/DeveloperDurp/durpot/model"
)

var (
	ApiKey = model.ApiKey
)

func HandleTag(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	client := openai.NewClient(ApiKey)

	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
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
