package score

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type client struct {
	scoreValue int
}

type Client interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	SetHandler(w http.ResponseWriter, r *http.Request)
	AddHandler(w http.ResponseWriter, r *http.Request)
}

func New(defaultValue int) Client {
	return &client{
		scoreValue: defaultValue,
	}
}

func (c *client) Get() int {
	return c.scoreValue
}

func (c *client) Set(val int) int {
	c.scoreValue = val
	return c.scoreValue
}

func (c *client) Add(val int) int {
	c.scoreValue += val
	return c.scoreValue
}

func (c *client) GetHandler(w http.ResponseWriter, r *http.Request) {
	s := c.Get()

	WriteJSONResponse(w, map[string]int{
		"value": s,
	})
}

type requestFormat struct {
	Value int `json:"value"`
}

func (c *client) SetHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rf requestFormat
	err := decoder.Decode(&rf)

	if err != nil {
		log.Errorf("decoder failed to decode: %v", err)
		WriteErrorResponse(w, err)
		return
	}

	val := rf.Value
	s := c.Set(val)

	WriteJSONResponse(w, map[string]int{
		"value": s,
	})
}

func (c *client) AddHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rf requestFormat
	err := decoder.Decode(&rf)

	if err != nil {
		log.Errorf("decoder failed to decode: %v", err)
		WriteErrorResponse(w, err)
		return
	}

	val := rf.Value
	s := c.Add(val)

	WriteJSONResponse(w, map[string]int{
		"value": s,
	})
}

// WriteErrorResponse writes an error back from an invalid request
func WriteErrorResponse(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// WriteJSONResponse writes some value and encodes into the response
func WriteJSONResponse(w http.ResponseWriter, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(val)
}
