package interfaces

import "fmt"

func (i Interfaces) DeleteInterface(secondaryExit bool) {
	var option int
	var id uint

	for !secondaryExit {
		fmt.Println("Delete")
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
			fmt.Println("ID informasi makanan yang ingin dihapus:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Konfirmasi:")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				i.S.DeleteFoodByID(id)
			case 2:
			}
		case 2:
			fmt.Println("ID informasi minuman yang ingin dihapus:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Konfirmasi:")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				i.S.DeleteDrinkByID(id)
			case 2:
			}
		case 3:
			fmt.Println("ID informasi staff yang ingin dihapus:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Konfirmasi:")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				i.S.DeleteStaffByID(id)
			case 2:
			}
		case 4:
			fmt.Println("ID informasi topping yang ingin dihapus:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Konfirmasi:")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				i.S.DeleteToppingByID(id)
			case 2:
			}
		case 9:
			secondaryExit = !secondaryExit
		}
		fmt.Println()

	}
}
