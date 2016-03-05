package config

import (
	"github.com/kuriouslabs/godo/models"
)

type Env struct {
	TodoRepo *models.TodoRepo
}

//TODO: Probably should pass in some sort of settings file or something
func MakeEnv() *Env {
	adaptor := models.NewInMemoryAdaptor()
	return &Env{
		TodoRepo: models.NewTodoRepo(adaptor),
	}
}
