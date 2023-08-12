package slogger

import (
	"context"
	"time"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

// LogData описывает структуру единицы логирования
type LogData struct {
	Level  slog.Level
	Attrs  []slog.Attr
}

// On создает единицу логирования с уровнем lvl
func On(lvl slog.Level) *LogData {
	return &LogData{
		Level: lvl,
	}
}

// Print добавляет данные в файл лога
func (ld *LogData) Print(msg string, args ...interface{}) {
	msg = strings.TrimSuffix(color.CyanString(fmt.Sprintln(append([]interface{}{msg}, args...)...)), "\n")

	switch ld.Level {
	case Debug:
		slog.Log(context.Background(), ld.Level, msg, ld.Attrs)
	case Notice:
		slog.Log(context.Background(), ld.Level, msg, ld.Attrs)
	case Info:
		slog.Info(msg, ld.Attrs)
	case Warn:
		slog.Warn(msg, ld.Attrs)
	case Fatal:
		os.Exit(1)
	default:
		slog.Info("Undefined log level - " + ld.Level.String())
	}
}

// wrapLogParams инициализирует параметры логгера
func wrapLogParams() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: false,
		Level:     &Lvler,
		ReplaceAttr: func(gps []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.TimeKey { // Форматирование временной метки
				attr.Value = slog.StringValue(fmt.Sprint(time.Now().Format("2006-01-02 15:04:05")))
			}

			if attr.Key == slog.LevelKey {
				// Вывод наименования уровня логирования
				level := attr.Value.Any().(slog.Level)
				levelLabel, exists := Levels[level]
				if !exists {
					levelLabel = level.String()
				}

				// Замена цвета уровня логирования
				f := ColorLevels[level]
				levelLabel = f(levelLabel)

				attr.Value = slog.StringValue(levelLabel)
			}

			return attr
		},
	}
}
