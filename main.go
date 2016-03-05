package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/config"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s!\n", ps.ByName("name"))
}

func main() {
	fmt.Println("Starting on port 5000")
    env := config.MakeEnv()
    fmt.Println(env)
    
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":5000", router))
}
