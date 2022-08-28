package service

import "Week4/model"

func (s Service) CreateFood(name string, price uint) {
	s.R.CreateFood(model.Food{Name: name, Price: price})
}

func (s Service) ReadFoodByID(id uint) model.Food {
	return s.R.ReadFoodByID(id)
}

func (s Service) ReadFoodByName(name string) model.Food {
	return s.R.ReadFoodByName(name)
}

func (s Service) UpdateFoodNameByID(id uint, value string) {
	s.R.UpdateFoodNameByID(id, value)
}

func (s Service) UpdateFoodPriceByID(id uint, value uint) {
	s.R.UpdateFoodPriceByID(id, value)
}

func (s Service) DeleteFoodByID(id uint) {
	s.R.DeleteFoodByID(id)
}
