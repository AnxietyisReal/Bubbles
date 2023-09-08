package loaders

import (
	"encoding/json"
	"fmt"
	"os"
)

type Tokens struct {
	Bot string `json:"bot"`
}

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

func LoadBotToken(path string) string {
	var tokens Tokens
	err := LoadJSON(path, &tokens)
	if err != nil {
		return "Bot token not found; Please create a \"tokens.json\" file in root directory and insert bot token in it."
	}
	return tokens.Bot
}
