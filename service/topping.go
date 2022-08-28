package service

import "Week4/model"

func (s Service) CreateTopping(topping string) {
	s.R.CreateTopping(model.Topping{Name: topping})
}

func (s Service) ReadToppingByID(id uint) model.Topping {
	return s.R.ReadToppingByID(id)
}

func (s Service) ReadToppingByName(name string) model.Topping {
	return s.R.ReadToppingByName(name)
}

func (s Service) UpdateToppingNameByID(id uint, value string) {
	s.R.UpdateToppingNameByID(id, value)
}

func (s Service) DeleteToppingByID(id uint) {
	s.R.DeleteToppingByID(id)
}
