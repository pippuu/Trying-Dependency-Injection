package interfaces

import (
	"Week4/pkg"
	"Week4/service"
	"fmt"
)

type Interfaces struct {
	S service.Service
}

func InitInterfaces(service service.Service) Interfaces {
	return Interfaces{
		S: service,
	}
}

func (i Interfaces) MenuInterfaces() {
	var option int
	var mainExit, secondaryExit bool

	mainExit = false
	for !mainExit {
		secondaryExit = false
		pkg.ClearTerminal()
		fmt.Println("MAIN MENU")
		fmt.Println()
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("9. Exit")
		fmt.Println()
		fmt.Println("YOUR OPTION:")
		fmt.Scanln(&option)
		fmt.Println()
		switch option {
		case 1:
			i.CreateInterface(secondaryExit)
		case 2:
			i.ReadInterface(secondaryExit)
		case 3:
			i.UpdateInterface(secondaryExit)
		case 4:

		case 9:
			mainExit = !mainExit
		}
	}
}
