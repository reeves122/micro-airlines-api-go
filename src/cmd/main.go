package main

import (
	"github.com/reeves122/micro-airlines-api-go/repository"
	"github.com/reeves122/micro-airlines-api-go/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	repo, err := repository.NewRepository("repo.db")
	if err != nil {
		panic(err)
	}

	log.Info("Starting server")
	svr := server.NewServer(repo)
	svr.Run(":3000")
}
