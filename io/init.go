package io

import "fmt"

func init() {
	//let get our server host list from docker container
	loadServers()

	//let get also our database name from docker container
	loadDB()

	//let let connect our cassandra cluster
	connect()
}

func SayHello() {
	fmt.Println("(:)-> io Cassandra connected")
}
