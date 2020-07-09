package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewTest(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) TestHello(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World Test")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello %s", d)
}