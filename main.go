package main

import (
	"github.com/ShuichiroTsuda/cloud/db"
	"github.com/ShuichiroTsuda/cloud/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
