package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/controllers"
	"github.com/kuriouslabs/godo/middleware"
)

type Router struct {
	router *httprouter.Router
	Render *render.Render
}

func NewRouter() *Router {
	router := &Router{
		router: httprouter.New(),
		Render: render.New(render.Options{}),
	}
	router.registerRoutes()
	router.registerAuthRoutes()
	return router
}

func (r *Router) registerRoutes() {
	// LogIn
	auth := controllers.NewAuthController()
	r.router.POST("/login", r.wrap(auth.LogIn))

	// User
	user := controllers.NewUserController()
	r.router.GET("/me", r.wrap(user.Me))
	r.router.POST("/user/create", r.wrap(user.Create))
}

func (r *Router) registerAuthRoutes() {
	// Todos
	t := controllers.NewTodoController()
	r.router.GET("/todos/:id", r.wrapAuth(t.Show))
	r.router.POST("/todos/create", r.wrapAuth(t.Create))

	user := controllers.NewUserController()
	r.router.GET("/user", r.wrapAuth(user.Me))
}

func (r *Router) wrap(h controllers.Action) httprouter.Handle {
	return middleware.Respond(r.Render, h)
}

func (r *Router) wrapAuth(h controllers.Action) httprouter.Handle {
	return middleware.Authenticated(r.wrap(h))
}

// ServeHTTP allows us to be a http.Handler
func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(rw, req)
}
