package data

import "testing"

func TestValidation(t *testing.T) {
	product := &Product{Name: "Tea"}
	err := product.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
