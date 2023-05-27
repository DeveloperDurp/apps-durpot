package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/DeveloperDurp/durpot/model"
)

var (
	ClientID = model.ClientID
	TokenURL = model.TokenURL
	Username = model.Username
	Password = model.Password
)

func GetUnraidUsage(s *discordgo.Session, m *discordgo.MessageCreate) {

	token := GenerateToken(ClientID, "client_credentials", TokenURL, Username, Password)

	url := "https://durpapi.durp.info/api/v1/unraid/powerusage"
	accessToken := token.AccessToken

	body := CallDurpAPI(url, accessToken)

	var response model.PowerUsageResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	jsonData, _ := json.MarshalIndent(response, "", "  ")
	jsonStr := string(jsonData)
	_, err = s.ChannelMessageSend(m.ChannelID, "Power Usage Response:\n```json\n"+jsonStr+"\n```")
	if err != nil {
		fmt.Println("Failed to send message")
	}

}

//func getSwansonQuote(s *discordgo.Session, m *discordgo.MessageCreate, url string) {
//	resp, err := http.Get(url + "/ronswanson")
//	if err != nil {
//		errStr := err.Error()
//		s.ChannelMessageSend(m.ChannelID, errStr)
//		return
//	}
//	defer resp.Body.Close()
//
//	var data []string
//	err = json.NewDecoder(resp.Body).Decode(&data)
//	if err != nil {
//		errStr := err.Error()
//		s.ChannelMessageSend(m.ChannelID, errStr)
//		return
//	}
//
//	if len(data) == 0 {
//		errStr := "No quotes found."
//		s.ChannelMessageSend(m.ChannelID, errStr)
//		return
//	}
//
//	s.ChannelMessageSend(m.ChannelID, data[0])
//}

//func sendAPIRequest(s *discordgo.Session, m *discordgo.MessageCreate, url string, endpoint string) {
//	var response interface{}
//	switch endpoint {
//	case "dadjoke":
//		url = url + "/dadjoke"
//		var data model.DadJokeResponse
//		response = &data
//	case "jinglebells":
//		url = url + "/foaas/jinglebells/durp"
//		var data model.JingleBellsResponse
//		response = &data
//	case "yomama":
//		url = url + "/yomama"
//		var data model.YomamaJokeResponse
//		response = &data
//	case "swanson":
//		getSwansonQuote(s, m, url)
//		return
//	default:
//		s.ChannelMessageSend(m.ChannelID, "Invalid endpoint.")
//		return
//	}
//
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		errStr := err.Error()
//		s.ChannelMessageSend(m.ChannelID, errStr)
//		return
//	}
//	req.Header.Set("Accept", "application/json")
//
//	client := http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		errStr := err.Error()
//		s.ChannelMessageSend(m.ChannelID, errStr)
//		return
//	}
//	defer resp.Body.Close()
//
//	err = json.NewDecoder(resp.Body).Decode(response)
//	if err != nil {
//		errStr := err.Error()
//		s.ChannelMessageSend(m.ChannelID, errStr)
//		return
//	}
//
//	switch endpoint {
//	case "dadjoke":
//		data := response.(*model.DadJokeResponse)
//		s.ChannelMessageSend(m.ChannelID, data.Joke)
//	case "jinglebells":
//		data := response.(*model.JingleBellsResponse)
//		s.ChannelMessageSend(m.ChannelID, data.Message)
//	case "yomama":
//		data := response.(*model.YomamaJokeResponse)
//		s.ChannelMessageSend(m.ChannelID, data.Joke)
//	}
//}
