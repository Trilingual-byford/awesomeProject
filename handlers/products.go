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
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	//catch All if no method is satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Oops" + r.Method))
}
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	marshal, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Write(marshal)
}
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	product := &data.Product{}
	err := product.GetFromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}
	addedProduct := data.AddProduct(product)
	toJson, err := addedProduct.MarshToJson()
	w.Write(toJson)
	p.l.Println("Prod:%#v", product)
}
