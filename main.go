package main

import (
	"fmt"

	"github.com/codegangsta/negroni"
	"github.com/kuriouslabs/godo/config"
	"github.com/kuriouslabs/godo/controllers"
)

func main() {
	fmt.Println("Starting on port 5000")

	env := config.MakeEnv()
	controllers.RegisterEnv(env)

	router := NewRouter()

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":5000")
}
