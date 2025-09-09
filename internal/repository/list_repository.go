package repository

import "trello-services/internal/entity"

type ListRepository interface {
	Create(list *entity.List) error
	GetByID(id string) (*entity.List, error)
	GetAll(page int, pageSize int) ([]entity.List, error)
	Update(list *entity.List) error
	Delete(id string) error
	Count() (int64, error)
}
