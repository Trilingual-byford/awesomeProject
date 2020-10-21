package handlers

import (
	"awesomeProject/data"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	marshal, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(marshal)
}
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	value := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&value)
	jsonProduct, err := value.MarshToJson()
	if err != nil {
		return
	}
	w.Write(jsonProduct)
	p.l.Println("Prod:%#v", value)
}
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p.l.Println("This is the Vars Id:", vars)

}

type KeyProduct struct{}

func (p Products) MiddleWareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		product := data.Product{}
		err := product.GetFromJson(request.Body)
		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			http.Error(writer, "Error Reading product", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(request.Context(), KeyProduct{}, product)
		validateErr := product.Validate()
		if validateErr != nil {
			http.Error(writer, "Validate error happened", http.StatusBadRequest)
			return
		}
		requestCtx := request.WithContext(ctx)
		//Call the next handler,which can be anther middleware in the chain,or the final handler
		next.ServeHTTP(writer, requestCtx)
	})

}
