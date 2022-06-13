package config

import "os"

type Configs struct {
	TargetHost   string
	ProxyHost    string
	Scheme       string
	SecretString string
	TargetSecret string
}

func GetConfigs() *Configs {
	return &Configs{
		os.Getenv("TARGET_HOST"),
		os.Getenv("PROXY_HOST"),
		os.Getenv("SCHEME"),
		os.Getenv("AUTH"),
		os.Getenv("TARGET_SECRET"),
	}
}
