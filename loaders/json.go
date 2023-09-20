package loaders

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tokens struct {
	DeployCommands bool   `json:"deployCommands"`
	Bot            string `json:"bot"`
	HookID         string `json:"hookId"`
	HookToken      string `json:"hookToken"`
	BotPublicKey   string `json:"botPublicKey"`
	Database       string `json:"database"`
	MongoUser      string `json:"mongoUser"`
	MongoPass      string `json:"mongoPass"`
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
	case "hookId":
		return tokens.HookID
	case "hookToken":
		return tokens.HookToken
	case "botPublicKey":
		return tokens.BotPublicKey
	case "database":
		return tokens.Database
	case "mongoUser":
		return tokens.MongoUser
	case "mongoPass":
		return tokens.MongoPass
	}
	return invalidToken
}

func IsCmdsDeployable() bool {
	err := LoadJSON("tokens.json", &tokens)
	if err != nil {
		return false
	}
	return tokens.DeployCommands
}
