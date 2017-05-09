package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/locona/ganesha/twitter"
)

type Config struct {
	Env string `envconfig:"APP_ENV" valid:"required"`
}

func main() {
	// scheduling every 15 min.
	twitter.GetSearch("birthday")
}

func LoadConfig() {
	var c Config
	envconfig.Process("", &c)
	godotenv.Load(fmt.Sprintf(".env.%s", c.Env))
}
