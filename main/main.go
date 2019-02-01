package main

import (
	service "github.com/mapleque/mpadmin"

	"github.com/mapleque/kelp/config"
	"github.com/mapleque/kelp/http"
	"github.com/mapleque/kelp/logger"
	"github.com/mapleque/kelp/mysql"
)

func main() {
	config.AddConfiger(config.ENV, "config", "")
	conf := config.Use("config")

	logdir := conf.Get("LOG_DIR") + conf.Get("HOSTNAME")
	if logdir != "" {
		logger.Add("http", logdir+"/http.log").SetTagOutput(logger.DEBUG, false)
	}

	log := logger.Get("http")
	mysql.SetLogger(log)
	http.SetLogger(log)

	// init database
	if err := mysql.AddDB(
		"database",
		conf.Get("DATABASE_DSN"),
		conf.Int("DATABASE_MAX_CONN"),
		conf.Int("DATABASE_MAX_IDLE"),
	); err != nil {
		log.Fatal("invalid database configure", err)
	}

	// init service
	conn := mysql.Get("database")
	ss := service.New(log, conn)

	// boot service
	go ss.Run(
		"0.0.0.0:"+conf.Get("HTTP"),
		conf.Get("SERVICE_HTTP_TOKEN"),
	)

	// waiting for exit
	select {}

}
