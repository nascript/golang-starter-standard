package wrapper

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

type ResponseFormatWrapper struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Page      int64  `json:"page"`
	TotalPage int64  `json:"total_page"`
	SortBy    string `json:"sort_by,omitempty"`
	SortDir   string `json:"sort_dir,omitempty"`
	Limit     int64  `json:"limit"`
}

type PaginationResponse[T any] struct {
	List       []*T       `json:"list"`
	Total      int64      `json:"total_data"`
	Pagination Pagination `json:"pagination"`
}

func (pr *PaginationResponse[T]) ParseData(ctx *fiber.Ctx) {
	pagination := Pagination{
		SortBy:  ctx.Params("sort_by"),
		SortDir: ctx.Params("sort_dir"),
	}
	defaultPage, defaultLimit := 1, 10
	pagination.Page = int64(defaultPage)
	if page := ctx.Query("page"); page != "" {
		if data, err := strconv.ParseInt(page, 0, 64); err == nil {
			pagination.Page = data
		}
	}
	pagination.Limit = int64(defaultLimit)
	if limit := ctx.Query("limit"); limit != "" {
		if data, err := strconv.ParseInt(limit, 0, 64); err == nil {
			pagination.Limit = data
		}
	}
	if sortBy := ctx.Query("sort_by"); sortBy != "" {
		pagination.SortBy = sortBy
	}
	if sortDir := ctx.Query("sort_dir"); sortDir != "" {
		pagination.SortBy = sortDir
	}
	pagination.TotalPage = int64(math.Ceil(float64(pr.Total) / float64(pagination.Limit)))
	pr.Pagination = pagination
}
