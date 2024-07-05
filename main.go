package main

import (
	"fmt"

	"github.com/fabiano-amaral/my-app-go/logger"
	"github.com/fabiano-amaral/my-app-go/slack"
)

func main() {
	log := logger.InitLog()
	logger.PrintHello()
	log.Info("Hello, World!")
	fmt.Println("----")
	slack.ReceiveMessage()
	slack.SendMessage()
	fmt.Println("----")
}
