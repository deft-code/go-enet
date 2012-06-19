package main

import "./enet"

import "flag"
import "fmt"
import "log"
import "net"
import "time"

var server_address = flag.String("client_address", "localhost:9998", "The address the server is listening on.")

//var client_address = flag.String("server_address","localhost:9997","The address the client will listen on.")

func main() {
	err := enet.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer enet.Deinitialize()

	host, err := enet.CreateHost(nil, 1, 2, 0, 0)
	if err != nil {
		log.Fatalf("Failed to create host: '%s'", err)
	}
	defer host.Destroy()

	server_addr, err := net.ResolveUDPAddr("udp", *server_address)
	if err != nil {
		log.Fatalf("Invalid udp address: '%s'", err)
	}

	server, err := host.Connect(server_addr, 2, 42)
	if err != nil {
		log.Fatal(err)
	}

	event, err := host.Service(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	if event == nil || event.Type != enet.CONNECT {
		log.Fatal("Failed to connect to server", event)
	}

	for {
		fmt.Print("send> ")
		var input string = "junk"
		fmt.Scan(&input)

		if input == "quit" {
			server.Disconnect(42)
			for event, err = host.Service(3 * time.Second); event != nil; {
				switch event.Type {
				case enet.DISCONNECT:
					fmt.Printf("disconnected")
					return
				default:
					fmt.Println("discarding events during disconnect", event)
				}
			}
			server.DisconnectNow(42)
			return
		}

		payload := []byte(input)
		server.Send(0, payload, enet.RELIABLE)

		for event, err = host.Service(time.Second); true; {
			if err != nil {
				log.Fatal(err)
			}

			if event == nil {
				break
			}

			switch event.Type {
			case enet.CONNECT:
				fmt.Printf("Connection made, %v\n", event.Data)
			case enet.DISCONNECT:
				fmt.Printf("Disconnection, %v\n", event.Data)
				return
			case enet.RECEIVE:
				msg := string(event.Packet)
				fmt.Printf("message: %v", msg)
			default:
				log.Fatal("unkown event", event)
			}

         break
		}
	}
}
