package repository

import "Week4/model"

func (r Repository) CreateDrink(drink model.Drink) {
	if result := r.DB.Create(&drink); result.Error != nil {
		panic(result.Error)
	}
}

func (r Repository) ReadDrinkByID(id uint) model.Drink {
	var drink model.Drink

	if result := r.DB.First(&drink, "id = ?", id); result.Error != nil {
		panic(result.Error)
	}
	return drink
}

func (r Repository) ReadDrinkByName(name string) model.Drink {
	var drink model.Drink

	if result := r.DB.First(&drink, "name = ?", name); result.Error != nil {
		panic(result.Error)
	}
	return drink
}

func (r Repository) UpdateDrinkNameByID(id uint, value string) {
	drink := r.ReadDrinkByID(id)
	drink.Name = value
	r.DB.Save(&drink)
}

func (r Repository) UpdateDrinkPriceByID(id uint, value uint) {
	drink := r.ReadDrinkByID(id)
	drink.Price = value
	r.DB.Save(&drink)
}

func (r Repository) DeleteDrinkByID(id uint) {
	drink := r.ReadDrinkByID(id)
	r.DB.Delete(&drink)
}
