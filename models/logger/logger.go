package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func AddFileLogger(logger *logrus.Logger, fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatalf("An error occurred: \n\t%s", err.Error())
	}
	defer file.Close()

	FileHook := &FormatterHook{
		Writer: file,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
		Formatter: &FileFormatter{
			TimestampFormat: "01/02/2006 15:03:04",
		},
	}

	logger.AddHook(FileHook)
}

func New(fileName ...string) *logrus.Logger {
	newLogger := logrus.New()
	newLogger.SetOutput(io.Discard)
	newLogger.SetLevel(logrus.InfoLevel)

	StdoutHook := &FormatterHook{
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
		},
		Formatter: &StdoutFormatter{},
	}
	newLogger.AddHook(StdoutHook)

	StderrHook := &FormatterHook{
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
		Formatter: &StderrFormatter{
			TimestampFormat: "01/02/2006 15:04:05",
		},
	}
	newLogger.AddHook(StderrHook)

	return newLogger
}
