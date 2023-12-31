package main

import (
	"fmt"
	"os"

	"github.com/kyallanum/athena/cmd"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, "\nAn error occured: ", err)
		}
	}()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
