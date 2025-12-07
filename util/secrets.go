package util

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Secrets struct {
	UserToken string `json:"user_token"`
	BotToken  string `json:"bot_token"`
}

func GetSecrets() Secrets {
	secrets := Secrets{}
	data := GetTokenFile()
	err := json.Unmarshal(data, &secrets)
	if err != nil {
		slog.Error("Failed to parse secrets from the file.")
		slog.Debug(err.Error())
		os.Exit(1)
	}
	return secrets
}
