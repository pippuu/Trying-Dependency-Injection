package repository

import "Week4/model"

func (r Repository) CreateFood(food model.Food) {
	if result := r.DB.Create(&food); result.Error != nil {
		panic(result.Error)
	}
}

func (r Repository) ReadFoodByID(id uint) model.Food {
	var food model.Food
	if result := r.DB.First(&food, "id = ?", id); result.Error != nil {
		panic(result.Error)
	}
	return food
}

func (r Repository) ReadFoodByName(name string) model.Food {
	var food model.Food
	if result := r.DB.First(&food, "name = ?", name); result.Error != nil {
		panic(result.Error)
	}
	return food
}

func (r Repository) UpdateFoodNameByID(id uint, value string) {
	food := r.ReadFoodByID(id)
	food.Name = value
	r.DB.Save(&food)
}

func (r Repository) UpdateFoodPriceByID(id uint, value uint) {
	food := r.ReadFoodByID(id)
	food.Price = value
	r.DB.Save(&food)
}

func (r Repository) DeleteFoodByID(id uint) {
	food := r.ReadFoodByID(id)
	r.DB.Delete(&food)
}
