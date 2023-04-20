package main

import (
	"log"

	"github.com/dany0814/go-apisolutions/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
