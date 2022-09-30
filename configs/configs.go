package configs

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type Configs struct {
	Appconfig AppConfig
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	if configs == nil {
		lock.Lock()

		if err := godotenv.Load(); err != nil {
			panic("Failed to load env file")
		}

		configs = &Configs{
			Appconfig: AppConfig{
				Name: os.Getenv("APP_NAME"),
				Env:  os.Getenv("APP_ENV"),
				Port: os.Getenv("APP_PORT"),
			},
		}
		lock.Unlock()
	}

	return configs
}
