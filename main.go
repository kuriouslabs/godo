package main

import (
	"fmt"

	"github.com/codegangsta/negroni"
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

	register := func(h middleware.DataHandler) httprouter.Handle {
		return middleware.Respond(r, h)
	}

	registerWithAuth := func(h middleware.DataHandler) httprouter.Handle {
		return middleware.Authenticated(register(h))
	}

	// LogIn
	auth := controllers.NewAuthController()
	router.POST("/login", register(auth.LogIn))

	// Todos
	t := controllers.NewTodoController(env.TodoRepo)
	router.GET("/todos/:id", registerWithAuth(t.Show))
	router.POST("/todos/create", registerWithAuth(t.Create))

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":5000")
}
