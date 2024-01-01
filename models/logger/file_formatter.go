package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type FileFormatter struct {
	logrus.TextFormatter
	TimestampFormat string
}

func (formatter *FileFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(formatter.TimestampFormat)
	return []byte(fmt.Sprintf("%s %s: %s\n", timestamp, entry.Level.String(), entry.Message)), nil
}
