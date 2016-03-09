package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/config"
	"github.com/kuriouslabs/godo/controllers"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	fmt.Println("Starting on port 5000")
	env := config.MakeEnv()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", controllers.Todo(env.TodoRepo).Show())

	log.Fatal(http.ListenAndServe(":5000", router))
}
