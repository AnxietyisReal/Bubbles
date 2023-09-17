package loaders

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tokens struct {
	Bot          string `json:"bot"`
	Webhook      string `json:"webhook"`
	BotPublicKey string `json:"botPublicKey"`
	Database     string `json:"database"`
}

var (
	tokens       Tokens
	invalidToken = "Token not found; Please setup a \"tokens.json\" file in root directory."
)

func LoadJSON(path string, v interface{}) error {
	var file, err = os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func TokenLoader(token string) string {
	err := LoadJSON("tokens.json", &tokens)
	if err != nil {
		return invalidToken
	}
	switch token {
	case "bot":
		return tokens.Bot
	case "webhook":
		return tokens.Webhook
	case "database":
		return tokens.Database
	case "botPublicKey":
		return tokens.BotPublicKey
	}
	return invalidToken
}
