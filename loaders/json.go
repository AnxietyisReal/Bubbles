package loaders

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tokens struct {
	Bot          string `json:"bot"`
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
	switch token {
	case "bot":
		err := LoadJSON("tokens.json", &tokens)
		if err != nil {
			return invalidToken
		}
		return tokens.Bot
	case "database":
		err := LoadJSON("tokens.json", &tokens)
		if err != nil {
			return invalidToken
		}
		return tokens.Database
	case "botPublicKey":
		err := LoadJSON("tokens.json", &tokens)
		if err != nil {
			return invalidToken
		}
		return tokens.BotPublicKey
	}
	return invalidToken
}
