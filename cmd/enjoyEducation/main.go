package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"

	"os"
)

type config struct {
	Port      int    `envconfig:"PORT" required:"true"`
	DBConnStr string `envconfig:"DB_CONN_STR" required:"true"`
}

func main() {
	log := zerolog.New(os.Stdout).With().Logger()

	var conf config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

}
