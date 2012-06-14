package main

import "./enet"

import "flag"
import "log"
import "net"

var address = flag.String("address", "localhost:9998", "The address the server will listen on.")

func main() {
	err := enet.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer enet.Deinitialize()

	addr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		log.Fatalf("Invalid udp address: '%s'", err)
	}

	host, err := enet.CreateHost(addr, 2, 2, 0, 0)
	if err != nil {
		log.Fatalf("Failed to create host: '%s'", err)
	}
	defer host.Destroy()

}
