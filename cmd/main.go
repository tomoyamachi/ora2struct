package main

import (
	"log"
	"os"
)

func main() {
	app := NewApp("dev")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
