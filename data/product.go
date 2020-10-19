package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

func GetProducts() []*Product {
	return productList
}
func (p *Product) GetFromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}
func (p *Product) MarshToJson() ([]byte, error) {
	marshal, err := json.Marshal(p)
	return marshal, err
}
func AddProduct(p *Product) *Product {
	p.ID = getNextID()
	productList = append(productList, p)
	printProduct(productList)
	return p
}
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
func printProduct(products []*Product) {
	fmt.Printf("len=%d cap=%d %v\n", len(products), cap(products), products)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc232",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       2.45,
		SKU:         "abc232sssss",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
