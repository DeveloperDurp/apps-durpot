package model

import "os"

var (
	Token     = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOTPREFIX")
	ChannelID = os.Getenv("ChannelID")
	ApiKey    = os.Getenv("OPENAI_API_KEY")
	ClientID  = os.Getenv("ClientID")
	TokenURL  = os.Getenv("TokenURL")
	Username  = os.Getenv("Username")
	Password  = os.Getenv("Password")
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

type PowerUsageResponse struct {
	Voltage12VLoad  int     `json:"12v_load"`
	Voltage12VWatts float64 `json:"12v_watts"`
	Voltage3VLoad   int     `json:"3v_load"`
	Voltage3VWatts  float64 `json:"3v_watts"`
	Voltage5VLoad   int     `json:"5v_load"`
	Voltage5VWatts  float64 `json:"5v_watts"`
	Capacity        string  `json:"capacity"`
	Efficiency      float64 `json:"efficiency"`
	FanRPM          int     `json:"fan_rpm"`
	Load            int     `json:"load"`
	PoweredOn       string  `json:"poweredon"`
	PoweredOnRaw    string  `json:"poweredon_raw"`
	Product         string  `json:"product"`
	Temperature1    float64 `json:"temp1"`
	Temperature2    float64 `json:"temp2"`
	Uptime          string  `json:"uptime"`
	UptimeRaw       string  `json:"uptime_raw"`
	Vendor          string  `json:"vendor"`
	TotalWatts      float64 `json:"watts"`
}
