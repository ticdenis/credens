package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error", err)
	}

	var reply struct {
		Text string
	}

	err = client.Call("API.SayHello", "world", &reply)
	if err != nil {
		log.Fatal("error calling API.SayHello", err)
	}

	log.Println(reply.Text)
}
