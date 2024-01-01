package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type FormatterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
	Formatter logrus.Formatter
}

func (hook *FormatterHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *FormatterHook) Levels() []logrus.Level {
	return hook.LogLevels
}
