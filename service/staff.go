package service

import "Week4/model"

func (s Service) CreateStaff(name string, age uint, address string) {
	s.R.CreateStaff(model.Staff{Name: name, Age: age, Address: address})
}

func (s Service) ReadStaffByID(id uint) model.Staff {
	return s.R.ReadStaffByID(id)
}

func (s Service) ReadStaffByName(name string) model.Staff {
	return s.R.ReadStaffByName(name)
}

func (s Service) UpdateStaffNameByID(id uint, value string) {
	s.R.UpdateStaffNameByID(id, value)
}

func (s Service) UpdateStaffAgeByID(id uint, value uint) {
	s.R.UpdateStaffAgeByID(id, value)
}

func (s Service) UpdateStaffAddressByID(id uint, value string) {
	s.R.UpdateStaffAddressByID(id, value)
}

func (s Service) DeleteStaffByID(id uint) {
	s.R.DeleteStaffByID(id)
}
