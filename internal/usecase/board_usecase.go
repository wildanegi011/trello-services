package usecase

import (
	"trello-services/internal/entity"
	"trello-services/internal/repository"
)

type BoardUsecase struct {
	repo repository.BoardRepository
}

func NewBoardUsecase(r repository.BoardRepository) *BoardUsecase {
	return &BoardUsecase{repo: r}
}

func (uc *BoardUsecase) Create(board *entity.Board) error {
	return uc.repo.Create(board)
}

func (uc *BoardUsecase) GetByID(id int) (*entity.Board, error) {
	return uc.repo.GetByID(id)
}

func (uc *BoardUsecase) GetAll(page int, pageSize int) ([]entity.Board, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	page = (page - 1) * pageSize
	boards, err := uc.repo.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := uc.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return boards, total, nil
}

func (uc *BoardUsecase) Update(board *entity.Board) error {
	return uc.repo.Update(board)
}

func (uc *BoardUsecase) Delete(id int) error {
	return uc.repo.Delete(id)
}
