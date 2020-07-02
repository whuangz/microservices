package main

import (
	"log"
	"net/http"
	"os"

	"gihtub.com/whuangz/product-api/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handler.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	log.Fatal(http.ListenAndServe(":8888", sm))
}
