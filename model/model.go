package model

type ConfigStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	ChannelID string `json:"ChannelID"`
	ApiKey    string `json:"OPEN_API_KEY"`
}

type JingleBellsResponse struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

type DadJokeResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type YomamaJokeResponse struct {
	Joke string `json:"joke"`
}
