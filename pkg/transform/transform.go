package transform

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Vector struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func NewVector(X float32, Y float32, Z float32) Vector {
	return Vector{X: X, Y: Y, Z: Z}
}

type client struct {
	pos Vector
	rot Vector
}

type Client interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	SetHandler(w http.ResponseWriter, r *http.Request)
}

func New(defaultPos Vector, defaultRot Vector) Client {
	return &client{
		pos: defaultPos,
		rot: defaultRot,
	}
}

func (c *client) Get() (Vector, Vector) {
	return c.pos, c.rot
}

func (c *client) Set(pos Vector, rot Vector) (Vector, Vector) {
	c.pos = pos
	c.rot = rot

	return c.pos, c.rot
}

func (c *client) GetHandler(w http.ResponseWriter, r *http.Request) {
	pos, rot := c.Get()

	WriteJSONResponse(w, map[string]Vector{
		"pos": pos,
		"rot": rot,
	})
}

type requestFormat struct {
	Pos Vector `json:"pos"`
	Rot Vector `json:"rot"`
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

	po := rf.Pos
	ro := rf.Rot
	pos, rot := c.Set(po, ro)

	WriteJSONResponse(w, map[string]Vector{
		"pos": pos,
		"rot": rot,
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
