package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOTPREFIX")

	return nil

}

var BotId string
var goBot *discordgo.Session

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

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	var response string
	baseurl := "https://kong.durp.info/"
	switch m.Content {
	case BotPrefix + "ping":
		response = "pong"
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "meme":
		response = getJson(baseurl+"random-meme", "url")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "catfact":
		response = getJson(baseurl+"cat-facts/fact", "fact")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "cat":
		response = getJson(baseurl+"random-cats", "file")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "yomama":
		response = getJson(baseurl+"yomama", "joke")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "dadjoke":
		response = getJson(baseurl+"dadjoke", "joke")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "dog":
		response = getJson(baseurl+"random-dogs", "message")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "jinglebells":
		response = getJson(baseurl+"foaas/jinglebells/durp", "message")
		s.ChannelMessageSend(m.ChannelID, response)
	case BotPrefix + "swanson":
		response = getJson(baseurl+"ronswanson", "base")
		s.ChannelMessageSend(m.ChannelID, response)
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

func getJson(url, value string) (output string) {

	var s string
	switch value {
	case "base":
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sc := strings.Trim(string(body), "[")
		sc = strings.Trim(sc, "]")
		sc = strings.Trim(sc, "\"")

		s = fmt.Sprintf("%v", sc)
	default:
		client := http.Client{
			Timeout: time.Second * 2, // Timeout after 2 seconds
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var result map[string]interface{}
		json.Unmarshal([]byte(b), &result)
		s = fmt.Sprintf("%v", result[value])
	}
	return string(s)
}
