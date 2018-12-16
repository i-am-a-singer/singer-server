package main

import (
	"os"

	"github.com/singerapi/singer-server/database"
	"github.com/singerapi/singer-server/rawdata"
	"github.com/singerapi/singer-server/service"
	flag "github.com/spf13/pflag"
)

// constant
const (
	PORT string = "8080"
	PATH string = "rawdata/rawdata.txt"
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
	server.Run(":" + port)

}
