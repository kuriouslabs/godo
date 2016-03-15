package main

import (
	"fmt"

	"github.com/codegangsta/negroni"
	"github.com/kuriouslabs/godo/config"
)

func main() {
	fmt.Println("Starting on port 5000")

	env := config.MakeEnv()
	router := NewRouter(env)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":5000")
}
