package service

import "Week4/model"

func (s Service) CreateFood(name string, price uint) {
	s.R.CreateFood(model.Food{Name: name, Price: price})
}

func (s Service) ReadFood(id uint) model.Food {
	return s.R.ReadFood(id)
}
