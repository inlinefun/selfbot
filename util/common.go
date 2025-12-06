package util

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type RunArgs struct {
	DebugMode bool
}

func Parse_args() RunArgs {
	args := RunArgs{}
	flag.BoolVar(&args.DebugMode, "debug", false, "Use debug mode")
	flag.Parse()
	return args
}

func Setup_logger(useDebugLevel bool) {

	level := slog.LevelInfo
	if useDebugLevel {
		level = slog.LevelDebug
	}

	writer := os.Stdout
	handler := tint.NewHandler(
		writer,
		&tint.Options{
			Level:      level,
			TimeFormat: time.Kitchen,
		},
	)
	slog.SetDefault(slog.New(handler))
}
