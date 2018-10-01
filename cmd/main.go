package main

import (
	"github.com/kevinwubert/go-simple-score/pkg/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := server.Main()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorf("Server errored with err: %v", err)
	}
}
