package values

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type client struct {
	m map[string]float32
}

type Client interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	SetHandler(w http.ResponseWriter, r *http.Request)
}

func New() Client {
	return &client{
		m: make(map[string]float32),
	}
}

func (c *client) Get(k string) float32 {
	return c.m[k]
}

func (c *client) Set(k string, v float32) float32 {
	c.m[k] = v

	return c.m[k]
}

func (c *client) GetHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]
	val := c.Get(key)

	WriteJSONResponse(w, map[string]float32{
		"val": val,
	})
}

type requestFormat struct {
	Key string  `json:"key"`
	Val float32 `json:"val"`
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

	v := c.Set(rf.Key, rf.Val)

	WriteJSONResponse(w, map[string]float32{
		"Val": v,
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
