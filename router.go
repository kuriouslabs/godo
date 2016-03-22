package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/controllers"
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
	r.router.POST("/token/refresh", r.wrap(auth.RefreshToken))

	// User
	user := controllers.NewUserController()
	r.router.POST("/user/create", r.wrap(user.Create))
}

func (r *Router) registerAuthRoutes() {
	// Todos
	t := controllers.NewTodoController()
	r.router.GET("/todos/:id", r.wrapAuth(t.Show))
	r.router.POST("/todos/create", r.wrapAuth(t.Create))

	auth := controllers.NewAuthController()
	r.router.POST("/logout", r.wrapAuth(auth.LogOut))

	user := controllers.NewUserController()
	r.router.GET("/user", r.wrapAuth(user.Me))
}

func (r *Router) wrap(h controllers.Action) httprouter.Handle {
	return controllers.Render(r.Render, h)
}

func (r *Router) wrapAuth(h controllers.Action) httprouter.Handle {
	return r.wrap(controllers.Authenticated(h))
}

// ServeHTTP allows us to be a http.Handler
func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(rw, req)
}
