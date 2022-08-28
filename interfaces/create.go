package interfaces

import (
	"Week4/model"
	"fmt"
)

func (i Interfaces) CreateInterface(secondaryExit bool) {
	var name, topping, address string
	var age, price uint
	var option int

	for !secondaryExit {
		fmt.Println("CREATE")
		fmt.Println()
		fmt.Println("1. Food")
		fmt.Println("2. Drink")
		fmt.Println("3. Staff")
		fmt.Println("4. Topping")
		fmt.Println("9. Exit")
		fmt.Println()
		fmt.Println("YOUR OPTION:")
		fmt.Scanln(&option)
		fmt.Println()
		switch option {
		case 1:
			fmt.Println("Nama:")
			fmt.Scanln(&name)
			fmt.Println("Harga:")
			fmt.Scanln(&price)
			i.S.CreateFood(name, price)
		case 2:
			fmt.Println("Nama:")
			fmt.Scanln(&name)
			fmt.Println("Harga:")
			fmt.Scanln(&price)
			i.S.CreateDrink(name, []*model.Topping{}, price)
		case 3:
			fmt.Println("Nama:")
			fmt.Scanln(&name)
			fmt.Println("Umur:")
			fmt.Scanln(&age)
			fmt.Println("Alamat:")
			fmt.Scanln(&address)
			i.S.CreateStaff(name, age, address)
		case 4:
			fmt.Println("Topping:")
			fmt.Scanln(&topping)
			i.S.CreateTopping(topping)
		case 9:
			secondaryExit = !secondaryExit
		}
		fmt.Println()

	}
}
