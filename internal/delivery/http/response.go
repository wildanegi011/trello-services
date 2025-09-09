package http

import "github.com/gin-gonic/gin"

type Metadata struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type Response struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Metadata *Metadata   `json:"metadata,omitempty"`
	Error    string      `json:"error,omitempty"`
}

func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Response{
		Status:   "success",
		Data:     data,
		Metadata: nil,
	})
}

func JsonSuccessWithMetadata(ctx *gin.Context, data interface{}, total int64, page, pageSize int) {
	metadata := &Metadata{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}
	ctx.JSON(200, Response{
		Status:   "success",
		Data:     data,
		Metadata: metadata,
	})
}

func JsonError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, Response{
		Status: "error",
		Error:  message,
	})
}
