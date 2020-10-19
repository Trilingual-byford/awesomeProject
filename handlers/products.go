package handlers

import (
	"awesomeProject/data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.Method==http.MethodGet{
		p.getProducts(w,r)
	}
	//catch All if no method is satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Oops"+r.Method))
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	marshal, err := json.Marshal(products)
	if err!=nil{
		http.Error(w,"Unable to marshal json",http.StatusInternalServerError)
	}
	w.Write(marshal)
}
