package main

import (
	"log"
	"os"

	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/rawdata"
	"github.com/singerapi/singer-server/service"
	flag "github.com/spf13/pflag"
)

// constant
var (
	PORT = "8080"
)

func main() {
	// create the bolt database
	database.OpenDB(rawdata.LoadRawData)
	defer database.Close()

	var port string
	port = os.Getenv("PORT")
	if len(port) <= 0 {
		port = PORT
	}

	flag.StringVarP(&port, "port", "p", "8080", "the port to listen")
	flag.Parse()

	server := service.NewServer()
	log.Printf("Client Home is served at: http://127.0.0.1:%s/singerapi/\n", port)
	server.Run(":" + port)

}
