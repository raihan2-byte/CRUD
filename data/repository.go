package data

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Nama, error)
	FindById(ID int) (Nama, error)
	Create(data Nama) (Nama, error)
	Update(data Nama) (Nama, error)
	Delete(data Nama) (Nama, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Nama, error) {
	var nama []Nama
	err := r.db.Find(&nama).Error

	return nama, err
}

func (r *repository) FindById(ID int) (Nama, error) {
	var nama Nama
	err := r.db.Find(&nama, ID).Error

	return nama, err
}

func (r *repository) Create(data Nama) (Nama, error) {
	err := r.db.Create(&data).Error

	return data, err
}

func (r *repository) Update(data Nama) (Nama, error) {
	err := r.db.Save(&data).Error

	return data, err
}

func (r *repository) Delete(data Nama) (Nama, error) {
	err := r.db.Delete(&data).Error

	return data, err
}
