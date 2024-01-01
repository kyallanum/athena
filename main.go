package main

import (
	"github.com/kyallanum/athena/cmd"
	"github.com/kyallanum/athena/models/logger"
)

func main() {
	logger := logger.New()

	defer func() {
		if err := recover(); err != nil {
			logger.Fatalf("\nAn error occured: \n\t%s", err)
		}
	}()

	if err := cmd.Execute(logger); err != nil {
		panic(err)
	}
}
