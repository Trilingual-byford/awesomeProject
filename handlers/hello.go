package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("start write Response")
	//writer.Write([]byte("Hi I am return"))
	d, err := ioutil.ReadAll(request.Body)
	http.Error(writer, "Oops", http.StatusBadRequest)
	//return
	if err != nil {
		//writer.Write([]byte("Hi I entered  nil"))
		http.Error(writer, "Oops", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s", d)}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

