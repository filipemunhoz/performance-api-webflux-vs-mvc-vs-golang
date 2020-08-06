package handlers

import (
	"log"
	"net/http"
	"performance-api-webflux-vs-mvc-vs-golang/performance-go/data"
	"strconv"
	"time"
)

// Products ...
type Products struct {
	l *log.Logger
}

// NewProducts ...
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	//p.l.Println("Handle GET Product")

	delay, _ := strconv.ParseInt(r.URL.Query().Get("delay"), 10, 32)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Error converting json", http.StatusInternalServerError)
	}
}
