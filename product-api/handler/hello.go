package handler

import (
    "fmt"
	"log"
    "net/http"
    "io/ioutil"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h * Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World Test")
    d, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Oops", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "hello %s", d)
}