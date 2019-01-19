package server

import (
	"net/http"

	"github.com/kevinwubert/go-simple-score/pkg/score"
	"github.com/kevinwubert/go-simple-score/pkg/transform"

	log "github.com/sirupsen/logrus"
)

func Main() error {
	log.Info("Starting simple score server...")

	initScore := 0
	scoreClient := score.New(initScore)

	initPos := transform.NewVector(0, 0, 0)
	initRot := transform.NewVector(0, 0, 0)

	transformClient := transform.New(initPos, initRot)

	initPos2 := transform.NewVector(0, 0, 0)
	initRot2 := transform.NewVector(0, 0, 0)

	transformClient2 := transform.New(initPos2, initRot2)

	http.HandleFunc("/getScore", scoreClient.GetHandler)
	http.HandleFunc("/setScore", scoreClient.SetHandler)
	http.HandleFunc("/addScore", scoreClient.AddHandler)

	http.HandleFunc("/getTransform", transformClient.GetHandler)
	http.HandleFunc("/setTransform", transformClient.SetHandler)
	http.HandleFunc("/getTransform2", transformClient2.GetHandler)
	http.HandleFunc("/setTransform2", transformClient2.SetHandler)

	err := http.ListenAndServe(":80", nil)
	return err
}
