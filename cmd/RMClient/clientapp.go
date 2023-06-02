package RMClient

import "RetailMonitor/config"

type ClientApp struct {
	cfg *config.Config
}

func Start(cfg *config.Config) (ClientApp, error) {

	return ClientApp{}, nil
}
