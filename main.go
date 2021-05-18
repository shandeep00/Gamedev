package main

import (
	"GetMega/Entities"
	"fmt"
	"strconv"
)

func main() {
	server := Entities.New()
	server.AddClient("p1")
	server.AddClient("p2")
	server.AddClient("p3")
	server.CreateGame("g1", 3, 6)
	server.ConnectGame(1, 1)
	server.ConnectGame(2, 1)
	server.ConnectGame(3, 1)
	for i := 1; i <= 6; i++ {
		server.AddState(1, strconv.Itoa(i))
		if i % 2 == 0 {
			server.UpdateClient(1, 2, i)
		} else {
			server.UpdateClient(1, 1, i)
		}

		if i == 3 {
			server.UpdateClient(2, 3, i)
		}

		if i % 4 == 0 {
			server.UpdateClient(3, 4, i)
		}
		for _, j := range server.GetClients(1) {
			fmt.Println("Client ", j, "received state : ", server.SendState(j, i))
		}
		fmt.Println("------------------------------------------")
	}
}