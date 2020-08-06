package main

import (
	"fmt"
	"net/rpc"
	"log"
)

type Item struct {
	Id string
	Plat string
	Pay string
}

func main(){
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"001", "|| F45672AA ||" , " Rp 40.000"}
	b := Item{"002", "|| A45678DS ||" , " Rp 60.000"}
	c := Item{"003", "|| B41238OP ||" , " Rp 30.000"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", Item{"002", "|| A45678DS ||" , " edit parking price: Rp 100.000"}, &reply)
	
	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)






}