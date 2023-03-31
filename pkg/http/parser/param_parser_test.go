package parser_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"skilledin-green-skills-api/pkg/http/parser"
	"testing"
)

func Test_ParsePaginationQueryParam(t *testing.T) {
	app := fiber.New()
	defer func() { _ = app.Shutdown() }()
	app.Get("/", func(ctx *fiber.Ctx) error {
		limit, offset := parser.ParsePaginationQueryParam(ctx)
		assert.Equal(t, limit, 10)
		assert.Equal(t, offset, 10)
		return nil
	})
	req := httptest.NewRequest("GET", "/?limit=10&page=2", nil)
	_, _ = app.Test(req)
}
