package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jasonmccallister/process"
)

func main() {
	// start one process
	go func() {
		if err := process.Start(process.Opts{
			Name:      "echo",
			Args:      []string{"Hello"},
			Writer:    os.Stdout,
			ErrWriter: os.Stderr,
		}); err != nil {
			log.Fatal(err)
		}

		if err := process.Start(process.Opts{
			Name:      "echo",
			Args:      []string{"world"},
			Writer:    os.Stdout,
			ErrWriter: os.Stderr,
		}); err != nil {
			log.Fatal(err)
		}
	}()

	// another blocking call such as a web server
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
