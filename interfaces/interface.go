package interfaces

import (
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
	var name string
	var price uint

	mainExit = false
	for !mainExit {
		secondaryExit = false
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
		switch option {
		case 1:
			for !secondaryExit {
				fmt.Println("CREATE")
				fmt.Println()
				fmt.Println("1. Food")
				fmt.Println("2. Drink")
				fmt.Println("3. Staff")
				fmt.Println("9. Exit")
				fmt.Println()
				fmt.Println("YOUR OPTION:")
				fmt.Scanln(&option)
				switch option {
				case 1:
					fmt.Println("Name:")
					fmt.Scanln(&name)
					fmt.Println("Price:")
					fmt.Scanln(&price)
					i.S.CreateFood(name, price)
				case 2:
				case 3:
				case 9:
					secondaryExit = !secondaryExit
				}
			}
		case 2:
			for !secondaryExit {
				fmt.Println("CREATE")
				fmt.Println()
				fmt.Println("1. Food")
				fmt.Println("2. Drink")
				fmt.Println("3. Staff")
				fmt.Println("9. Exit")
				fmt.Println()
				fmt.Println("YOUR OPTION:")
				fmt.Scanln(&option)
				switch option {
				case 1:
					fmt.Println("Name:")
					fmt.Scanln(&name)
					fmt.Println("Price:")
					fmt.Scanln(&price)
					i.S.CreateFood(name, price)
				case 2:
				case 3:
				case 9:
					secondaryExit = !secondaryExit
				}
			}
		case 9:
			mainExit = !mainExit
		}
	}

}
