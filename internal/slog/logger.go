package slog

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(env, path string) *Logger {
	var level slog.Level

	switch env {
	case EnvLocal:
		level = slog.LevelDebug
	case EnvDev:
		level = slog.LevelDebug
	case EnvProd:
		level = slog.LevelInfo
	default:
		panic(errors.New("invalid environment level"))
	}

	handlerOptions := &slog.HandlerOptions{
		Level:     level,
		AddSource: false,
	}

	//err := os.MkdirAll("logs", 0770)
	//if err != nil {
	//	panic(err)
	//}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	// defer f.Close()

	w := io.MultiWriter(os.Stderr, f)

	h := slog.NewTextHandler(w, handlerOptions)
	return &Logger{slog.New(h)}
	// slog.SetDefault(slog.New(h))
}

func (l *Logger) Infof(template string, args ...any) {
	l.Logger.Info(fmt.Sprintf(template, args))
}

func (l *Logger) Info(args ...any) {
	l.Logger.Info(fmt.Sprint(args))
}

func (l *Logger) Debugf(template string, args ...any) {
	l.Logger.Debug(fmt.Sprintf(template, args))
}

func (l *Logger) Debug(args ...any) {
	l.Logger.Debug(fmt.Sprint(args))
}

func (l *Logger) Errorf(template string, args ...any) {
	l.Logger.Error(fmt.Sprintf(template, args))
}

func (l *Logger) Error(args ...any) {
	l.Logger.Error(fmt.Sprint(args))
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.Logger.Error(fmt.Sprintf(template, args))
	os.Exit(1)
}

func (l *Logger) Fatal(args ...any) {
	l.Logger.Error(fmt.Sprint(args))
	os.Exit(1)
}
