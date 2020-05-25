package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	app := NewApp("dev")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
