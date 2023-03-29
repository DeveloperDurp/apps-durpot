package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

var (
	Token     string
	BotPrefix string
	ChannelID string
	BotId     string
	apiKey    string
	goBot     *discordgo.Session

	config *configStruct
)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
	ChannelID string `json : "ChannelID"`
}

type jingleBellsResponse struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

type dadJokeResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type yomamaJokeResponse struct {
	Joke string `json:"joke"`
}

func ReadConfig() error {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOTPREFIX")
	ChannelID = os.Getenv("ChannelID")
	apiKey = os.Getenv("OPENAI_API_KEY")

	return nil

}

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
	goBot.AddHandler(handleGuildMemberAdd)
	goBot.AddHandler(handleGuildMemberRemove)
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

	baseurl := "https://kong.durp.info/"
	switch m.Content {
	case BotPrefix + "ping":
		s.ChannelMessageSend(m.ChannelID, "pong")
	//case BotPrefix + "meme":
	//	response = getJson(baseurl+"random-meme", "url")
	//	s.ChannelMessageSend(m.ChannelID, response)
	//case BotPrefix + "catfact":
	//	response = getJson(baseurl+"cat-facts/fact", "fact")
	//	s.ChannelMessageSend(m.ChannelID, response)
	//case BotPrefix + "cat":
	//	response = getJson(baseurl+"random-cats", "file")
	//	s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "yomama":
		sendAPIRequest(s, m, baseurl, "yomama")
	case BotPrefix + "dadjoke":
		sendAPIRequest(s, m, baseurl, "dadjoke")
	//case BotPrefix + "dog":
	//	response = getJson(baseurl+"random-dogs", "message")
	//	s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "jinglebells":
		sendAPIRequest(s, m, baseurl, "jinglebells")
	case BotPrefix + "swanson":
		sendAPIRequest(s, m, baseurl, "swanson")

	}

}

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
	return
}

func getSwansonQuote(s *discordgo.Session, m *discordgo.MessageCreate, url string) {
	resp, err := http.Get(url + "/ronswanson")
	if err != nil {
		errStr := err.Error()
		s.ChannelMessageSend(m.ChannelID, errStr)
		return
	}
	defer resp.Body.Close()

	var data []string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		errStr := err.Error()
		s.ChannelMessageSend(m.ChannelID, errStr)
		return
	}

	if len(data) == 0 {
		errStr := "No quotes found."
		s.ChannelMessageSend(m.ChannelID, errStr)
		return
	}

	s.ChannelMessageSend(m.ChannelID, data[0])
}

func sendAPIRequest(s *discordgo.Session, m *discordgo.MessageCreate, url string, endpoint string) {
	var response interface{}
	switch endpoint {
	case "dadjoke":
		url = url + "/dadjoke"
		var data dadJokeResponse
		response = &data
	case "jinglebells":
		url = url + "/foaas/jinglebells/durp"
		var data jingleBellsResponse
		response = &data
	case "yomama":
		url = url + "/yomama"
		var data yomamaJokeResponse
		response = &data
	case "swanson":
		getSwansonQuote(s, m, url)
		return
	default:
		s.ChannelMessageSend(m.ChannelID, "Invalid endpoint.")
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errStr := err.Error()
		s.ChannelMessageSend(m.ChannelID, errStr)
		return
	}
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errStr := err.Error()
		s.ChannelMessageSend(m.ChannelID, errStr)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		errStr := err.Error()
		s.ChannelMessageSend(m.ChannelID, errStr)
		return
	}

	switch endpoint {
	case "dadjoke":
		data := response.(*dadJokeResponse)
		s.ChannelMessageSend(m.ChannelID, data.Joke)
	case "jinglebells":
		data := response.(*jingleBellsResponse)
		s.ChannelMessageSend(m.ChannelID, data.Message)
	case "yomama":
		data := response.(*yomamaJokeResponse)
		s.ChannelMessageSend(m.ChannelID, data.Joke)
	}
}

func handleGuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	message := fmt.Sprintf("Welcome <@%s> to our server!", m.Member.User.ID)
	_, err := s.ChannelMessageSend(ChannelID, message)
	if err != nil {
		log.Printf("Error sending welcome message: %v\n", err)
	}
}

func handleGuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	message := fmt.Sprintf("Goodbye %s", m.Member.User.Username)
	_, err := s.ChannelMessageSend(ChannelID, message)
	if err != nil {
		log.Printf("Error sending goodbye message: %v\n", err)
	}
}

func handleTag(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	client := openai.NewClient(apiKey)

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
