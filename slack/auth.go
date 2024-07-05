package slack

import "fmt"

var token string

func init() {
	token = "TESTE_123"
	fmt.Println("Authenticated with Slack")
}
