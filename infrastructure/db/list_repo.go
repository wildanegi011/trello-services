package db

import (
	"trello-services/internal/entity"
	"trello-services/internal/repository"

	"gorm.io/gorm"
)

type ListRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) repository.ListRepository {
	return &ListRepo{db: db}
}

func (r *ListRepo) Create(list *entity.List) error {
	return r.db.Create(list).Error
}

func (r *ListRepo) GetByID(id string) (*entity.List, error) {
	var list entity.List
	if err := r.db.Where("id = ?", id).First(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func (r ListRepo) GetAll(page int, page_size int) ([]entity.List, error) {
	var lists []entity.List
	if err := r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *ListRepo) Update(list *entity.List) error {
	return r.db.Save(list).Error
}

func (r *ListRepo) Delete(id string) error {
	return r.db.Delete(&entity.List{}, id).Error
}

func (r *ListRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entity.List{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
