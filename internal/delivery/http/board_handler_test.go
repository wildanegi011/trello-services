package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"trello-services/internal/entity"
	"trello-services/internal/usecase"

	"github.com/gin-gonic/gin"
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

func (m *MockBoardRepository) Count() (int64, error)                    { return 2, nil }
func (m *MockBoardRepository) Create(board *entity.Board) error         { return nil }
func (m *MockBoardRepository) GetByID(id string) (*entity.Board, error) { return nil, nil }
func (m *MockBoardRepository) Update(board *entity.Board) error         { return nil }
func (m *MockBoardRepository) Delete(id string) error                   { return nil }

func TestBoardHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := &MockBoardRepository{}
	boardUsecase := usecase.NewBoardUsecase(mockRepo)
	handler := &BoardHandler{uc: boardUsecase}
	router.GET("/api/v1/boards", handler.GetAll)

	req, _ := http.NewRequest("GET", "/api/v1/boards?page=1&page_size=2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Board 1")
	assert.Contains(t, w.Body.String(), "Board 2")
	assert.Contains(t, w.Body.String(), `"total":2`)
	assert.Contains(t, w.Body.String(), `"page":1`)
	assert.Contains(t, w.Body.String(), `"page_size":2`)
}
