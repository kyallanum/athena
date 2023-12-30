package main

import (
	"github.com/kyallanum/athena/v1.0.0/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		return
	}
}
