package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"performance-go/data"
	"strconv"
)

type Client struct {
	l *log.Logger
}

func NewClient(l *log.Logger) *Client {
	return &Client{l}
}

func (c *Client) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		delay, _ := strconv.ParseInt(r.URL.Query().Get("delay"), 10, 32)
		resp, err := http.Get(fmt.Sprintf("http://localhost:8083/products?delay=%d", delay))
		if err != nil {
			c.l.Print(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		var products []data.Product
		err = json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			c.l.Print(err)
		}
		err = json.NewEncoder(rw).Encode(products)
		if err != nil {
			c.l.Print(err)
		}
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
