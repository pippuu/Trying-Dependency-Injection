package repository

import "Week4/model"

func (r Repository) GetAllStaff() {
	var staffs []model.Staff
	r.DB.Find(&staffs)
}

func (r Repository) CreateStaff(staff model.Staff) {
	if result := r.DB.Create(&staff); result.Error != nil {
		panic(result.Error)
	}
}

func (r Repository) ReadStaffByID(id uint) model.Staff {
	var staff model.Staff

	if result := r.DB.First(&staff, "id = ?", id); result.Error != nil {
		panic(result.Error)
	}
	return staff
}

func (r Repository) ReadStaffByName(name string) model.Staff {
	var staff model.Staff

	if result := r.DB.First(&staff, "name = ?", name); result.Error != nil {
		panic(result.Error)
	}
	return staff
}

func (r Repository) UpdateStaffNameByID(id uint, value string) {
	staff := r.ReadStaffByID(id)
	staff.Name = value
	if result := r.DB.Save(&staff); result != nil {
		panic(result.Error)
	}
}

func (r Repository) UpdateStaffAgeByID(id uint, value uint) {
	staff := r.ReadStaffByID(id)
	staff.Age = value
	if result := r.DB.Save(&staff); result != nil {
		panic(result.Error)
	}
}

func (r Repository) UpdateStaffAddressByID(id uint, value string) {
	staff := r.ReadStaffByID(id)
	staff.Address = value
	if result := r.DB.Save(&staff); result != nil {
		panic(result.Error)
	}
}

func (r Repository) DeleteStaffByID(id uint) {
	staff := r.ReadStaffByID(id)
	if result := r.DB.Delete(&staff); result != nil {
		panic(result.Error)
	}
}
