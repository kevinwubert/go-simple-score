package score

import "net/http"

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

func (c *client) Set(val int) {
	c.scoreValue = val
}

func (c *client) Add(val int) {
	c.scoreValue += val
}

func (c *client) GetHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *client) SetHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *client) AddHandler(w http.ResponseWriter, r *http.Request) {

}
