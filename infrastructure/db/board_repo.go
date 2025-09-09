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

func (r *BoardRepo) GetByID(id string) (*entity.Board, error) {
	var board entity.Board
	if err := r.db.Where("id = ?", id).First(&board).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func (r BoardRepo) GetAll(page int, page_size int) ([]entity.Board, error) {
	var boards []entity.Board
	if err := r.db.
		Offset((page-1)*page_size).
		Limit(page_size).
		Preload("Lists", func(db *gorm.DB) *gorm.DB {
			return db.Order("position asc")
		}).
		Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

func (r *BoardRepo) Update(board *entity.Board) error {
	return r.db.Save(board).Error
}

func (r *BoardRepo) Delete(id string) error {
	return r.db.Delete(&entity.Board{}, id).Error
}

func (r *BoardRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Board{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
