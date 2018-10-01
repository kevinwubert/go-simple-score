package server

import (
	"net/http"

	"github.com/kevinwubert/go-simple-score/pkg/score"
	log "github.com/sirupsen/logrus"
)

func Main() error {
	log.Info("Starting simple score server...")

	initScore := 0
	scoreClient := score.New(initScore)

	http.HandleFunc("/getScore", scoreClient.GetHandler)
	http.HandleFunc("/setScore", scoreClient.SetHandler)
	http.HandleFunc("/addScore", scoreClient.AddHandler)

	err := http.ListenAndServe(":62802", nil)
	return err
}
