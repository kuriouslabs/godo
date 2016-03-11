package config

import "github.com/kuriouslabs/godo/repos"

type Env struct {
	TodoRepo *repos.TodoRepo
}

//TODO: Probably should pass in some sort of settings file or something
func MakeEnv() *Env {
	adaptor := repos.NewInMemoryAdaptor()
	return &Env{
		TodoRepo: repos.NewTodoRepo(adaptor),
	}
}
