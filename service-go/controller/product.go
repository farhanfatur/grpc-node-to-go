package controller

import (
	"log"
	"net/http"
)

type Product struct {
	L *log.Logger
}

func NewProduct(L *log.Logger) *Product {
	return &Product{
		L: L,
	}
}

func (p *Product) Login(w http.ResponseWriter, r *http.Request) {\
	p.L.Println("Hello World")
	w.Write([]byte("Hello World"))
}
