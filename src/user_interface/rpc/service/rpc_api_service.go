package service

import "fmt"

type RPCAPIService int

type Item struct {
	Text string
}

func (api *RPCAPIService) SayHello(to string, reply *Item) error {
	(*reply).Text = fmt.Sprintf("Hello %s from RPC app!", to)

	return nil
}
