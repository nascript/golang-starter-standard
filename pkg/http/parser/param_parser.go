package parser

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ParsePaginationQueryParam(ctx *fiber.Ctx) (limit, offset int) {
	defaultLimit := 10
	defaultPage := 1

	if limit := ctx.Query("limit"); limit != "" {
		defaultLimit, _ = strconv.Atoi(limit)
	}

	if page := ctx.Query("page"); page != "" {
		defaultPage, _ = strconv.Atoi(page)
	}

	defaultOffset := (defaultLimit * defaultPage) - defaultLimit

	return defaultLimit, defaultOffset
}
