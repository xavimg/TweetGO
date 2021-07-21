package main

import (
	"log"

	"github.com/xavimg/TweetGO/bd"
	"github.com/xavimg/TweetGO/handlers"
)

func main() {
	if bd.ChequeoConn() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
