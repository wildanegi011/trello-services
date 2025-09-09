package db

import (
	"trello-services/internal/entity"
	"trello-services/internal/repository"

	"gorm.io/gorm"
)

type BoardRepo struct {
	db *gorm.DB
}

func NewBoardRepo(db *gorm.DB) repository.BoardRepository {
	return &BoardRepo{db: db}
}

func (r *BoardRepo) Create(board *entity.Board) error {
	return r.db.Create(board).Error
}

func (r *BoardRepo) GetByID(id int) (*entity.Board, error) {
	var board entity.Board
	if err := r.db.First(&board, id).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func (r BoardRepo) GetAll(page int, page_size int) ([]entity.Board, error) {
	var boards []entity.Board
	if err := r.db.Offset((page - 1) * page_size).Limit(page_size).Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

func (r *BoardRepo) Update(board *entity.Board) error {
	return r.db.Save(board).Error
}

func (r *BoardRepo) Delete(id int) error {
	return r.db.Delete(&entity.Board{}, id).Error
}

func (r *BoardRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Board{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
