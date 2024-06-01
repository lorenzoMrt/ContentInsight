package main

import (
	"log"

	"github.com/lorenzoMrt/ContentInsight/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
