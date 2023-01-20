package main

import (
	"aliceServer/helpers"
	"fmt"
	"log"
	"os"
)

func main() {
	database := helpers.InitDatabase("alice", "quot-server", "3306", "alice", "AliceRoot1337")
	webApp := helpers.InitWebApp(database)

	// Listen on environment specified PORT, "3000" otherwise
	port := os.Getenv("PORT")
	if port == "" {
		port = "8005"
	}
	log.Fatal(webApp.Listen(fmt.Sprintf(":%s", port)))
}
