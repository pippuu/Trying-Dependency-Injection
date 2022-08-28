package interfaces

import "fmt"

func (i Interfaces) UpdateInterface(secondaryExit bool) {
	var option int
	var id uint
	var newName, newAddress string
	var newPrice, newAge uint

	for !secondaryExit {
		fmt.Println("UPDATE")
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
			fmt.Println("ID informasi makanan yang ingin diupdate:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Bagian yang ingin dirubah:")
			fmt.Println("1. Nama")
			fmt.Println("2. Harga")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("Nama makanan baru:")
				fmt.Scanln(&newName)
				i.S.UpdateFoodNameByID(id, newName)
			case 2:
				fmt.Println("Harga makanan baru:")
				fmt.Scanln(&newPrice)
				i.S.UpdateFoodPriceByID(id, newPrice)
			}
		case 2:
			fmt.Println("ID informasi minuman yang ingin diupdate:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Bagian yang ingin dirubah:")
			fmt.Println("1. Nama")
			fmt.Println("2. Harga")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("Nama minuman baru:")
				fmt.Scanln(&newName)
				i.S.UpdateDrinkNameByID(id, newName)
			case 2:
				fmt.Println("Harga minuman baru:")
				fmt.Scanln(&newPrice)
				i.S.UpdateDrinkPriceByID(id, newPrice)
			}
		case 3:
			fmt.Println("ID informasi staff yang ingin diupdate:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Bagian yang ingin dirubah:")
			fmt.Println("1. Nama")
			fmt.Println("2. Umur")
			fmt.Println("3. Alamat")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("Nama staff baru:")
				fmt.Scanln(&newName)
				i.S.UpdateStaffNameByID(id, newName)
			case 2:
				fmt.Println("Umur staff baru:")
				fmt.Scanln(&newAge)
				i.S.UpdateStaffAgeByID(id, newAge)
			case 3:
				fmt.Println("Alamat staff baru:")
				fmt.Scanln(&newAddress)
				i.S.UpdateStaffAddressByID(id, newAddress)
			}
		case 4:
			fmt.Println("ID informasi topping yang ingin diupdate:")
			fmt.Scanln(&id)
			fmt.Println()
			fmt.Println("Bagian yang ingin dirubah:")
			fmt.Println("1. Nama")
			fmt.Println()
			fmt.Println("YOUR OPTION:")
			fmt.Scanln(&option)
			fmt.Println()
			switch option {
			case 1:
				fmt.Println("Nama topping baru:")
				fmt.Scanln(&newName)
				i.S.UpdateToppingNameByID(id, newName)
			}
		case 9:
			secondaryExit = !secondaryExit
		}
		fmt.Println()

	}
}
