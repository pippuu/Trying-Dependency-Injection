package repository

import "Week4/model"

func (r Repository) CreateFood(food model.Food) {
	if result := r.DB.Create(&food); result.Error != nil {
		panic(result.Error)
	}
}

func (r Repository) ReadFood(id uint) model.Food {
	var food model.Food
	if err := r.DB.First(&food, "id = ?", id).Error; err != nil {
		panic(err.Error())
	}
	return food
}
