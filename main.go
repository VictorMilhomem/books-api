package main

import (
	"bookApi/database"
	"bookApi/server"
)

func main() {

	database.StartDB()

	serv := server.NewServer()

	serv.Run()
}
