package handlers

import (
	"log"
	"net/http"
	"performance-go/data"
	"strconv"
	"time"
)

// Products ...
type ProductsHandler struct {
	l *log.Logger
}

// NewProducts ...
func NewProducts(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (p *ProductsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *ProductsHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	delay, _ := strconv.ParseInt(r.URL.Query().Get("delay"), 10, 32)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Error converting json", http.StatusInternalServerError)
	}
}
