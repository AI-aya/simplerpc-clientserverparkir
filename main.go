package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)


type Item struct {
	Id string
	Plat string
	Pay string
}

type API int

var database []Item


func (a *API) GetDB(id string, reply *[]Item)error{
	*reply = database
	return nil
}


func (a *API) GetById(Id string, reply *Item) error{
	var getItem Item

	for _, val := range database {
		if val.Id == Id {
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item

	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Id == edit.Id {
			database[idx] = Item{edit.Id, edit.Plat, edit.Pay}
			changed = database[idx]
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Id == item.Id && val.Plat == item.Plat && val.Pay == item.Pay{
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}



func main(){

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)

	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	
	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}






}