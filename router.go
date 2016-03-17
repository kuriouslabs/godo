package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/config"
	"github.com/kuriouslabs/godo/controllers"
	"github.com/kuriouslabs/godo/middleware"
)

type Router struct {
	router *httprouter.Router
	env    *config.Env
}

func NewRouter(env *config.Env) *Router {
	router := &Router{
		router: httprouter.New(),
		env:    env,
	}
	router.registerRoutes()
	router.registerAuthRoutes()
	return router
}

func (r *Router) registerRoutes() {
	// LogIn
	auth := controllers.NewAuthController()
	r.router.POST("/login", r.wrap(auth.LogIn))
}

func (r *Router) registerAuthRoutes() {
	// Todos
	t := controllers.NewTodoController(r.env.TodoRepo)
	r.router.GET("/todos/:id", r.wrapAuth(t.Show))
	r.router.POST("/todos/create", r.wrapAuth(t.Create))
}

func (r *Router) wrap(h controllers.Action) httprouter.Handle {
	return middleware.Respond(r.env.Render, h)
}

func (r *Router) wrapAuth(h controllers.Action) httprouter.Handle {
	return middleware.Authenticated(r.wrap(h))
}

// ServeHTTP allows us to be a http.Handler
func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(rw, req)
}
