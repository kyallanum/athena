package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type StdoutFormatter struct {
	logrus.TextFormatter
}

func (formatter *StdoutFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}
