package main

import (
	"gitub.com/whuangz/product-api/product-api/handler"
    "log"
    "net/http"
    "os"
)

func main() {
    l := log.New(os.Stdout, "product-api", log.LstdFlags)
    hh := handler.NewHello(l)

    sm := http.NewServeMux()
    sm.Handle("/", hh)


    log.Fatal(http.ListenAndServe(":8888", sm))
}