package main

import (
	"github.com/inlinefun/selfbot/util"
)

func main() {
	args := util.ParseRunArgs()
	util.SetupLogger(args.DebugMode)
}
