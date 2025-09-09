package repository

import "trello-services/internal/entity"

type BoardRepository interface {
	Create(board *entity.Board) error
	GetByID(id string) (*entity.Board, error)
	GetAll(page int, pageSize int) ([]entity.Board, error)
	Update(board *entity.Board) error
	Delete(id string) error
	Count() (int64, error)
}
