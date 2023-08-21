package task

import (
	"gorm.io/gorm"
)

type (
	Repositories interface {
		GetAll() ([]Task, error)
		Create(name string) error
		Delete(id int) error
		Complete(id int) error
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repositories {
	return &repo{
		db: db,
	}
}

func (r repo) GetAll() ([]Task, error) {
	var t []Task
	result := r.db.Table("task").Model(&t).Find(&t)

	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (r repo) Create(name string) error {
	t := Task{
		Name:     name,
		Complete: "false",
	}
	result := r.db.Table("task").Create(&t)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r repo) Delete(id int) error {
	t := Task{}
	result := r.db.Table("task").Delete(&t, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r repo) Complete(id int) error {
	t := Task{}
	result := r.db.Table("task").Model(&t).Where("id = ?", id).Update("complete", "true")
	if result.Error != nil {
		return result.Error
	}

	return nil
}
