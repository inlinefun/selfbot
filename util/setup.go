package util

import (
	"log/slog"
	"os"
	"syscall"

	"golang.org/x/term"
)

func getSecretsWithInstructions() Secrets {
	secrets := Secrets{}
	slog.Info("Make 100% sure to not mix up your tokens as they will make the selfbot not work.")
	slog.Info("Do note that the token you enter will not be visible.")
	print("Enter your user account's token: ")
	secrets.UserToken = getSecureInput()
	println()
	print("Enter your bot account's token: ")
	secrets.BotToken = getSecureInput()
	return secrets
}

func getSecureInput() string {
	data, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		slog.Error("Error reading token from the std input.")
		slog.Debug(err.Error())
		os.Exit(1)
	}
	return string(data)
}
