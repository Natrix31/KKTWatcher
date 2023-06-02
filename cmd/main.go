package main

import (
	"RetailMonitor/cmd/RMClient"
	"RetailMonitor/config"
)

func main() {
	cfg := config.GetConfig()
	var client, err = RMClient.Start(cfg)
}
