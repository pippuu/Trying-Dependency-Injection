package repository

import "Week4/model"

func (r Repository) CreateTopping(topping model.Topping) {
	if result := r.DB.Create(&topping); result.Error != nil {
		panic(result.Error)
	}
}

func (r Repository) ReadToppingByID(id uint) model.Topping {
	var topping model.Topping
	if result := r.DB.First(&topping, "id = ?", id); result.Error != nil {
		panic(result.Error)
	}
	return topping
}

func (r Repository) ReadToppingByName(name string) model.Topping {
	var topping model.Topping
	if result := r.DB.First(&topping, "name = ?", name); result.Error != nil {
		panic(result.Error)
	}
	return topping
}

func (r Repository) UpdateToppingNameByID(id uint, value string) {
	topping := r.ReadToppingByID(id)
	topping.Name = value
	r.DB.Save(&topping)
}

func (r Repository) DeleteToppingByID(id uint) {
	topping := r.ReadToppingByID(id)
	r.DB.Delete(&topping)
}
