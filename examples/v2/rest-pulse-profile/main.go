package main

import (
	"log"

	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	c := rest.NewClient()

	profile, err := c.Pulse.PublicPulseProfile("Bitfinex")
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(profile)
}
