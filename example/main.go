package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

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
	fmt.Println("Accessory Info:")
	dump(ai)

	lights, err := c.GetLights(context.Background())
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Lights:")
	dump(lights)

	if len(lights) == 0 {
		return
	}

	fmt.Println("Turning light off ...")

	for _, light := range lights {
		light.On = elgatoring.BoolFrom(false)
	}

	lights, err = c.SetLights(context.Background(), lights)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Lights:")
	dump(lights)

	time.Sleep(5 * time.Second)

	fmt.Println("Ok, let's ask the light to identify itself ...")
	if err := c.Identify(context.Background()); err != nil {
		log.Panic(err)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Next, let's turn the light on ...")

	for _, light := range lights {
		light.On = elgatoring.BoolFrom(true)
	}

	lights, err = c.SetLights(context.Background(), lights)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Lights:")
	dump(lights)

	time.Sleep(5 * time.Second)

	fmt.Println("Ok, let's adjust the brightness ...")
	for i := 0; i <= 100; i += 10 {
		b := max(3, i)
		fmt.Printf("Setting brightness to %d\n", b)
		for _, light := range lights {
			light.Brightness = b
		}

		lights, err = c.SetLights(context.Background(), lights)
		if err != nil {
			log.Panic(err)
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("Oof, that's bright, let's tone it down ...")
	for _, light := range lights {
		light.Brightness = 20
	}

	lights, err = c.SetLights(context.Background(), lights)
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Now let's play with the temperature ...")
	for t := elgatoring.MinTemperature; t <= elgatoring.MaxTemperature; t += 20 {
		fmt.Printf("Setting temperature to %dK\n", t.Kelvin())
		for _, light := range lights {
			light.Temperature = t
		}

		lights, err = c.SetLights(context.Background(), lights)
		if err != nil {
			log.Panic(err)
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("Ok, we're done. Shut it down!")
	for _, light := range lights {
		light.On = elgatoring.BoolFrom(false)
	}

	lights, err = c.SetLights(context.Background(), lights)
	if err != nil {
		log.Panic(err)
	}
}
