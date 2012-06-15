package main

import "./enet"

import "flag"
import "log"
import "net"
import "time"

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

	for {
		event, err := host.Service(3 * time.Second)
		if err != nil {
			log.Fatal(err)
		}

		if event == nil {
			continue
		}

		switch event.Type {
		case enet.CONNECT:
			log.Log("new connection: ", event.Data)
		case enet.DISCONNECT:
			log.Log("disconnection: ", event.Data)
		case enet.RECEIVE:
			log.Log("received: ", event.Packet)
			msg = string(event.Packet)
			switch msg {
			case "stop":
				event.Peer.Disconnect(42)
			case "stopall":
				return
			case "die":
				return
			default:
				event.Peer.Send(0, event.Packet, enet.RELIABLE)
			}
		}
	}
}
