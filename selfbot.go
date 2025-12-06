package main

import (
	"log/slog"

	"github.com/inlinefun/selfbot/util"
)

func main() {
	args := util.Parse_args()
	util.Setup_logger(args.DebugMode)
	slog.Error("E")
	slog.Warn("W")
	slog.Info("I")
	slog.Debug("D")
}
