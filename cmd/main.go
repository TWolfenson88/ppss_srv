package main

import (
	"log"
	"net/http"
	"test_server/conf"
)

func main() {
	conf.CreateDBConfiguration()

	//here we should read some flags, conf, etc and then run CreateAndListen() func

	CreateAndListen()
}

//CreateAndListen function read server settings and launch server
func CreateAndListen() {
	server := &http.Server{}

	conf.Config(server)

	log.Println("Server is listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
