package config

import (
	"github.com/kuriouslabs/godo/repos"
	"gopkg.in/unrolled/render.v1"
)

type Env struct {
	TodoRepo *repos.TodoRepo
	Render   *render.Render
}

//TODO: Probably should pass in some sort of settings file or something
func MakeEnv() *Env {
	adaptor := repos.NewInMemoryAdaptor()
	return &Env{
		TodoRepo: repos.NewTodoRepo(adaptor),
		Render:   render.New(render.Options{}),
	}
}
