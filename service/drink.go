package service

import "Week4/model"

func (s Service) CreateDrink(name string, toppings []*model.Topping, price uint) {
	s.R.CreateDrink(model.Drink{Name: name, Topping: toppings, Price: price})
}

func (s Service) ReadDrinkByID(id uint) model.Drink {
	return s.R.ReadDrinkByID(id)
}

func (s Service) ReadDrinkByName(name string) model.Drink {
	return s.R.ReadDrinkByName(name)
}

func (s Service) UpdateDrinkNameByID(id uint, value string) {
	s.R.UpdateDrinkNameByID(id, value)
}

func (s Service) UpdateDrinkPriceByID(id uint, value uint) {
	s.R.UpdateDrinkPriceByID(id, value)
}

func (s Service) DeleteDrinkByID(id uint) {
	s.R.DeleteDrinkByID(id)
}
