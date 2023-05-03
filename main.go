package main

import (
	"wrap-midjourney/handlers"
	"wrap-midjourney/initialization"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	cfg := pflag.StringP("config", "c", "./config.yaml", "api server config file path.")

	pflag.Parse()

	initialization.LoadConfig(*cfg)
	initialization.LoadDiscordClient(handlers.DiscordMsgCreate, handlers.DiscordMsgUpdate)

	r := gin.Default()

	r.Use(handlers.CorsHandler())

	r.POST("/v1/trigger/midjourney-bot", handlers.MidjourneyBot)

	r.Run(":16007")
}
