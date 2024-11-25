package config

import "os"

type Config struct {
	DatabasePath string
	KeystorePath string
	NodeUrl      string
}

func Load() Config {
	return Config{
		DatabasePath: os.Getenv("DATABASE_PATH"),
		KeystorePath: os.Getenv("KEYSTORE_PATH"),
		NodeUrl:      os.Getenv("INFURA_URL"),
	}
}
