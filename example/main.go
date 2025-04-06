package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/kellegous/elgatoring"
)

func dump(data any) {
	b, err := json.MarshalIndent(data, " ", "  ")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s\n", b)
}

func main() {
	var flags struct {
		Host string
	}

	flag.StringVar(&flags.Host, "host", "127.0.0.1", "network address of the light")
	flag.Parse()

	c, err := elgatoring.New(flags.Host)
	if err != nil {
		log.Panic(err)
	}

	ai, err := c.GetAccessoryInfo(context.Background())
	if err != nil {
		log.Panic(err)
	}

	dump(ai)

	lights, err := c.GetLights(context.Background())
	if err != nil {
		log.Panic(err)
	}

	dump(lights)
}
