package android

import "github.com/netbirdio/netbird/client/internal"

func ReadConfig(configPath string) (*internal.Config, error) {
	return internal.ReadConfig(configPath)
}
