package interfaces

import "fmt"

func (i Interfaces) ReadInterface(secondaryExit bool) {
	var option int
	var id uint
	var name string

	for !secondaryExit {
		fmt.Println("FIND")
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
			fmt.Println("FIND BY")
			fmt.Println()
			fmt.Println("1. ID")
			fmt.Println("2. Nama")
			fmt.Println("9. EXIT")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("ID makanan yang ingin dicari:")
				fmt.Scanln(&id)
				food := i.S.ReadFoodByID(id)
				fmt.Println("Nama makanan:", food.Name)
				fmt.Println("Harga makanan:", food.Price)
			case 2:
				fmt.Println("Nama makanan yang ingin dicari:")
				fmt.Scanln(&name)
				food := i.S.ReadFoodByName(name)
				fmt.Println("Nama makanan:", food.Name)
				fmt.Println("Harga makanan:", food.Price)
			}
		case 2:
			fmt.Println("FIND BY")
			fmt.Println()
			fmt.Println("1. ID")
			fmt.Println("2. Nama")
			fmt.Println("9. EXIT")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("ID minuman yang ingin dicari:")
				fmt.Scanln(&id)
				drink := i.S.ReadDrinkByID(id)
				fmt.Println("Nama minuman:", drink.Name)
				fmt.Println("Topping minuman:", drink.Topping)
				fmt.Println("Harga minuman:", drink.Price)
			case 2:
				fmt.Println("Nama minuman yang ingin dicari:")
				fmt.Scanln(&name)
				drink := i.S.ReadDrinkByName(name)
				fmt.Println("Nama minuman:", drink.Name)
				fmt.Println("Topping minuman:", drink.Topping)
				fmt.Println("Harga minuman:", drink.Price)
			}
		case 3:
			fmt.Println("FIND BY")
			fmt.Println()
			fmt.Println("1. ID")
			fmt.Println("2. Nama")
			fmt.Println("9. EXIT")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("ID staff yang ingin dicari:")
				fmt.Scanln(&id)
				staff := i.S.ReadStaffByID(id)
				fmt.Println("Nama topping:", staff.Name)
				fmt.Println("Umur staff:", staff.Age)
				fmt.Println("Alamat staff:", staff.Address)
			case 2:
				fmt.Println("Nama staff yang ingin dicari:")
				fmt.Scanln(&name)
				staff := i.S.ReadStaffByName(name)
				fmt.Println("Nama staff:", staff.Name)
				fmt.Println("Umur staff:", staff.Age)
				fmt.Println("Alamat staff:", staff.Address)
			}
		case 4:
			fmt.Println("FIND BY")
			fmt.Println()
			fmt.Println("1. ID")
			fmt.Println("2. Nama")
			fmt.Println("9. EXIT")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("ID topping yang ingin dicari:")
				fmt.Scanln(&id)
				topping := i.S.ReadToppingByID(id)
				fmt.Println("Nama topping:", topping.Name)
			case 2:
				fmt.Println("Nama topping yang ingin dicari:")
				fmt.Scanln(&name)
				topping := i.S.ReadToppingByName(name)
				fmt.Println("Nama topping:", topping.Name)
			}
		case 9:
			secondaryExit = !secondaryExit
		}
		fmt.Println()

	}
}
