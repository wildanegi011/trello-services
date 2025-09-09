package repository

import "trello-services/internal/entity"

type BoardRepository interface {
	Create(board *entity.Board) error
	GetByID(id int) (*entity.Board, error)
	GetAll(page int, pageSize int) ([]entity.Board, error)
	Update(board *entity.Board) error
	Delete(id int) error
	Count() (int64, error)
}
