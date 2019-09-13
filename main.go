package main

import (
	"context"
	"fmt"
	"gitlab.com/NebulousLabs/go-upnp"
	"log"
)

func main() {
	device, err := upnp.DiscoverCtx(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ip, err := device.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("External IP address:", ip)

	err = device.Forward(9001, "go-upnp-forward-port")
	if err != nil {
		log.Fatal(err)
	}
}
