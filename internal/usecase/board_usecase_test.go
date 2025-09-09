package usecase_test

import (
	"testing"
	"trello-services/internal/entity"
	"trello-services/internal/usecase"

	"github.com/stretchr/testify/assert"
)

type MockBoardRepository struct{}

func (m *MockBoardRepository) GetAll(page int, pageSize int) ([]entity.Board, error) {
	boards := []entity.Board{
		{ID: "1", OrgID: "org1", Title: "Board 1"},
		{ID: "2", OrgID: "org2", Title: "Board 2"},
	}
	return boards, nil
}

func (m *MockBoardRepository) Count() (int64, error) {
	return 2, nil
}

func (m *MockBoardRepository) GetByID(id string) (*entity.Board, error) {
	return &entity.Board{ID: "1", OrgID: "org1", Title: "Board 1"}, nil
}

// Unused methods for interface compliance
func (m *MockBoardRepository) Create(board *entity.Board) error { return nil }
func (m *MockBoardRepository) Update(board *entity.Board) error { return nil }
func (m *MockBoardRepository) Delete(id string) error           { return nil }

func TestBoardUsecase_GetAll(t *testing.T) {
	mockRepo := &MockBoardRepository{}
	uc := usecase.NewBoardUsecase(mockRepo)

	boards, total, err := uc.GetAll(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, boards, 2)
	assert.Equal(t, "Board 1", boards[0].Title)
	assert.Equal(t, "Board 2", boards[1].Title)
}

func TestBoardUsecase_GetByID(t *testing.T) {
	mockRepo := &MockBoardRepository{}
	uc := usecase.NewBoardUsecase(mockRepo)

	board, err := uc.GetByID("1")
	assert.NoError(t, err)
	assert.NotNil(t, board)
	assert.Equal(t, "1", board.ID)
	assert.Equal(t, "org1", board.OrgID)
	assert.Equal(t, "Board 1", board.Title)
}
