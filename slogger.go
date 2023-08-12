package slogger

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

const (
	Debug  = slog.Level(-8)
	Notice = slog.Level(-4)
	Info   = slog.Level(0)
	Warn   = slog.Level(4)
	Fatal  = slog.Level(8)
)

// Levels описывает наименования произвольные уровни логирования
var Levels = map[slog.Leveler]string{
	Debug:  "[DEBUG]",
	Notice: "[NOTICE]",
	Info:   "[INFO]",
	Warn:   "[WARNING]",
	Fatal:  "[FATAL]",
}

// ColorLevels описывает цветовую схему произвольных уровней логирования
var ColorLevels = map[slog.Leveler]func(format string, a ...interface{}) string {
	Debug: color.BlueString,
	Notice: color.YellowString,
	Info: color.GreenString,
	Warn: color.RedString,
	Fatal: color.MagentaString,
}

// Lvler регулирует уровень логирования
var Lvler slog.LevelVar

// Make создает структурированный лог и заменяет лог по умолчанию
func Make(logPath string, defaultLevel slog.Level) {
	// Установка уровня логирования
	Lvler.Set(defaultLevel)

	var w io.Writer
	if logPath != "" {
		_, err := os.Stat(logPath)
		if os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
				os.Stderr.WriteString("не удалось создать директорию для сохранения файла лога")
			}
		
			err = ioutil.WriteFile(logPath, nil, os.ModePerm)
			if err != nil {
				os.Stderr.WriteString("не удалось записать файл лога")
			}
		}

		w, err = os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			w = os.Stdout
		}

		slog.New(slog.NewJSONHandler(w, wrapLogParams()))
	} else {
		slog.New(slog.NewTextHandler(os.Stdout, wrapLogParams()))
	}
}
