package config

import "github.com/kuriouslabs/godo/repos"

type Env struct {
	TodoRepo  *repos.TodoRepo
	UserRepo  *repos.UserRepo
	TokenRepo *repos.TokenRepo
}

//TODO: Probably should pass in some sort of settings file or something
func MakeEnv() *Env {
	adaptor := repos.NewInMemoryAdaptor()
	return &Env{
		TodoRepo:  repos.NewTodoRepo(adaptor),
		UserRepo:  repos.NewUserRepo(adaptor),
		TokenRepo: repos.NewTokenRepo(),
	}
}
