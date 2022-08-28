package main

import (
	"Week4/config"
	"Week4/interfaces"
	"Week4/repository"
	"Week4/service"
)

func main() {
	// Init for config
	C := config.InitConfig()
	// Init for repo
	R := repository.InitRepository(C)
	// Init for service
	S := service.InitService(R)
	// Init for interface
	I := interfaces.InitInterfaces(S)
	I.MenuInterfaces()
}
