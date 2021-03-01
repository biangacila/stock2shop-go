package io

import "os"

func loadServers() {
	if os.Getenv("HOST1") != "" {
		Servers = append(Servers, os.Getenv("HOST1"))
	}
	if os.Getenv("HOST2") != "" {
		Servers = append(Servers, os.Getenv("HOST2"))
	}
	if os.Getenv("HOST3") != "" {
		Servers = append(Servers, os.Getenv("HOST3"))
	}
	if os.Getenv("HOST4") != "" {
		Servers = append(Servers, os.Getenv("HOST4"))
	}
	if len(Servers) == 0 {
		//Servers = append(Servers, "uzuri.easipath.com")
		Servers = append(Servers, "server1.easipath.com")
		Servers = append(Servers, "server2.easipath.com")
		Servers = append(Servers, "service3.easipath.com")
	}
}

func loadDB() {
	if os.Getenv("DB_NAME") != "" {
		DB_NAME = os.Getenv("DB_NAME")
	} else {
		DB_NAME = "demo"
	}
}
