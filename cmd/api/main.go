package main

import (
	"https://github.com/lorenzoMrt/ContentInsight/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
