package main

import "github.com/sweetnebulae/go_ent/config"

func main() {
	client := config.ConnectDB()
	defer config.DisconnectDB(client)

	select {}
}
