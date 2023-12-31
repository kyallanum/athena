package main

import (
	"log"
	"os"

	"github.com/kyallanum/athena/v1.0.0/cmd"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	defer func() {
		if err := recover(); err != nil {
			logger.Fatalln("\nAn issue occurred: ", err)
		}
	}()

	if err := cmd.Execute(logger); err != nil {
	  panic(err)	
	}
}
