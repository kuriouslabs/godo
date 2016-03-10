package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/config"
	"github.com/kuriouslabs/godo/controllers"
	"github.com/kuriouslabs/godo/middleware"
	"gopkg.in/unrolled/render.v1"
)

func main() {
	fmt.Println("Starting on port 5000")
	env := config.MakeEnv()

	router := httprouter.New()
	r := render.New(render.Options{})

	withAuth := middleware.Authenticated

	// Todos
	t := controllers.NewTodoController(env.TodoRepo)
	router.GET("/todos/:id", withAuth(middleware.Respond(r, t.Show)))
	router.POST("/todos/create", withAuth(middleware.Respond(r, t.Create)))

	log.Fatal(http.ListenAndServe(":5000", router))
}
