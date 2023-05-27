package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.com/DeveloperDurp/durpot/model"
)

func CallDurpAPI(url string, accesstoken string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+accesstoken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}
	return body
}

func GenerateToken(clientID string, grantType string, url string, username string, password string) model.AccessTokenResponse {

	formData := fmt.Sprintf("grant_type=%s&client_id=%s&username=%s&password=%s",
		grantType, clientID, username, password)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(formData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return model.AccessTokenResponse{}
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return model.AccessTokenResponse{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return model.AccessTokenResponse{}
	}

	var response model.AccessTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return model.AccessTokenResponse{}
	}

	return response
}
