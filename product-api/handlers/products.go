package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/whuangz/product-api/domains"
)

type Products struct {
	l *log.Logger
}

func NewProducs(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(w, req)
		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(w, req)
		return
	}

	if req.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(req.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.addProduct(id, w, req)

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, req *http.Request) {
	lp := domains.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, req *http.Request) {

	prod := &domains.Product{}
	err := prod.FromJSON(req.Body)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusBadRequest)
	}

	domains.AddProduct(prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, req *http.Request) {

	prod := &domains.Product{}
	err := prod.FromJSON(req.Body)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusBadRequest)
	}

	err := domains.UpdateProduct(id, prod)
	if err == domains.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
