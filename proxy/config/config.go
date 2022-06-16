package config

import "os"

type Configs struct {
	TargetHost   string
	ProxyHost    string
	Scheme       string
	ProxySecret  string
	TargetSecret string
}

func GetConfigs() *Configs {
	return &Configs{
		os.Getenv("TARGET_HOST"),
		os.Getenv("PROXY_HOST"),
		os.Getenv("SCHEME"),
		os.Getenv("PROXY_SECRET"),
		os.Getenv("TARGET_SECRET"),
	}
}
