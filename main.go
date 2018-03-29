package main

import (
	"log"

	"github.com/t0m/code_resume/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
