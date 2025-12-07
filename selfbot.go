package main

import (
	"github.com/inlinefun/selfbot/client"
	"github.com/inlinefun/selfbot/util"
)

func main() {
	args := util.ParseRunArgs()
	util.SetupLogger(args.DebugMode)
	secrets := util.GetSecrets()
	client.NewBotClient(secrets)
	client.NewUserClient(secrets)
}
