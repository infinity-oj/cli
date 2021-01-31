package main

import (
	"log"
	"os"
)

func main() {
	app, err := CreateApp()
	if err != nil {
		panic(err)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
