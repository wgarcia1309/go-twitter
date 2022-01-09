package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wgarcia1309/go-twitter/db"
	"github.com/wgarcia1309/go-twitter/routers"
)

func main() {

	if !db.CheckConecction() {
		log.Fatal("no db connection")
	}
	routers.Routers()
}
