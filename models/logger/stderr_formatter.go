package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type StderrFormatter struct {
	logrus.Formatter
	TimestampFormat string
}

func (formatter *StderrFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(formatter.TimestampFormat)
	return []byte(fmt.Sprintf("%s Error: %s\n", timestamp, entry.Message)), nil
}
