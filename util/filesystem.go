package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

const configSubPath = "selfbot"
const configFile = "config.json"
const secretsFile = "secrets.json"

func getConfigDir() string {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		slog.Error("Error accessing the user's config directory.")
		slog.Debug(err.Error())
		os.Exit(1)
	}
	dir := filepath.Join(configRoot, configSubPath)
read:
	_, err = os.ReadDir(dir)
	switch {
	case err == nil:
		return dir
	case errors.Is(err, fs.ErrNotExist):
		slog.Debug(err.Error())
		slog.Info(fmt.Sprintf("The config directory does not exist, creating one at '%s'", dir))
		os.Mkdir(dir, fs.ModePerm.Perm())
		goto read
	case errors.Is(err, fs.ErrPermission):
		slog.Error("Failed to access the config directory due to missing permissions.")
		slog.Debug(err.Error())
		os.Exit(1)
	default:
		slog.Error("Unhandled error reading the config directory.")
		slog.Debug(err.Error())
		os.Exit(1)
	}
	return dir
}

func GetTokenFile() []byte {
	dir := getConfigDir()
	file := filepath.Join(dir, secretsFile)
read:
	data, err := os.ReadFile(file)
	switch {
	case err == nil:
		return data
	case errors.Is(err, fs.ErrNotExist):
		secrets := getSecretsWithInstructions()
		json, err := json.MarshalIndent(secrets, "", "    ")
		if err != nil {
			slog.Error("Error serializing secrets to a json.")
			slog.Debug(err.Error())
		}
		WriteConfigFile(secretsFile, json)
		goto read
	case errors.Is(err, fs.ErrPermission):
		slog.Error("Failed to read the secrets file due to missing permissions.")
		slog.Debug(err.Error())
		os.Exit(1)
	default:
		slog.Error(fmt.Sprintf("Unhandled error while reading %s", file))
		slog.Debug(err.Error())
		os.Exit(1)
	}
	return data
}

func WriteConfigFile(filename string, data []byte) {
	dir := getConfigDir()
	file := filepath.Join(dir, filename)
	err := os.WriteFile(file, data, os.ModePerm.Perm())
	switch {
	case err == nil:
		return
	case errors.Is(err, fs.ErrPermission):
		slog.Error(fmt.Sprintf("Failed to write %s at %s due to missing permissions.", filename, configSubPath))
		slog.Debug(err.Error())
		os.Exit(1)
	default:
		slog.Error(fmt.Sprintf("Unhandled error while writing %s at %s", filename, configSubPath))
		slog.Debug(err.Error())
		os.Exit(1)
	}
}
