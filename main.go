package main

import (
	"github.com/mheers/nats-seeder/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	// execeute the command
	err := cmd.Execute()
	if err != nil {
		logrus.Fatalf("Execute failed: %+v", err)
	}
}
