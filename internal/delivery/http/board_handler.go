package http

import (
	"net/http"
	"strconv"
	"trello-services/internal/entity"
	"trello-services/internal/usecase"

	"github.com/gin-gonic/gin"
)

type BoardHandler struct {
	uc *usecase.BoardUsecase
}

func NewBoardHandler(r *gin.RouterGroup, uc *usecase.BoardUsecase) {
	h := &BoardHandler{uc}

	r.POST("/boards", h.Create)
	r.GET("/boards", h.GetAll)
	r.GET("/boards/:id", h.GetByID)
}

func (h *BoardHandler) Create(ctx *gin.Context) {
	var board entity.Board
	if err := ctx.ShouldBindJSON(&board); err != nil {
		JsonError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.uc.Create(&board); err != nil {
		JsonError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	JsonSuccess(ctx, board)
}

func (h *BoardHandler) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	boards, total, err := h.uc.GetAll(page, pageSize)

	if err != nil {
		JsonError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	JsonSuccessWithMetadata(ctx, boards, total, page, pageSize)
}

func (h *BoardHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	board, err := h.uc.GetByID(id)
	if err != nil {
		JsonError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	JsonSuccess(ctx, board)
}
