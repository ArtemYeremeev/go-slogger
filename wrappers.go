package slogger

import (
	"strconv"
	"runtime"

	"golang.org/x/exp/slog"
)

// WithSource оборачивает источника ошибки
func (ld *LogData) WithSource() *LogData {
	_, path, line, ok := runtime.Caller(1)
	if !ok {
		return ld
	}

	a := slog.Attr{
		Key:   "source",
		Value: slog.StringValue(path + ":" + strconv.Itoa(line)),
	}
	ld.Attrs = append(ld.Attrs, a)

	return ld
}
