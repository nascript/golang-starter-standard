package rest_test

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"skilledin-green-skills-api/mocks"
	"skilledin-green-skills-api/modules/skill/delivery/rest"
	"skilledin-green-skills-api/modules/skill/domain/response"
	"skilledin-green-skills-api/pkg/http/wrapper"
	"testing"
)

func Test_Skill_CoverageSetupSuite(_ *testing.T) {
	app := fiber.New()
	rest.NewGreenSkillRESTHandler(app, new(mocks.ISkillService))
}

func Test_Skill_Fetch_ShouldSuccess(t *testing.T) {
	app := fiber.New()
	svcMock := new(mocks.ISkillService)
	svcMock.On("GreenSkillList",
		mock.Anything, mock.Anything,
		mock.Anything, mock.Anything).Return(
		&wrapper.PaginationResponse[response.GreenSkillResponse]{
			List:  []*response.GreenSkillResponse{{ID: "123"}},
			Total: 1,
		}, nil).Once()
	path := "/api/v1/green-skills"
	handler := rest.GreenSkillRESTHandler{Service: svcMock}
	app.Get(path, func(ctx *fiber.Ctx) error {
		ctx.Query("search", "lorem")
		ctx.Query("limit", "1")
		ctx.Query("page", "1")
		_ = ctx.Next()
		return nil
	}, handler.Fetch)
	req := httptest.NewRequest("GET", path, nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.StatusCode, fiber.StatusOK)
}
func Test_Skill_Fetch_ShouldError(t *testing.T) {
	app := fiber.New()
	svcMock := new(mocks.ISkillService)
	svcMock.On("GreenSkillList",
		mock.Anything, mock.Anything,
		mock.Anything, mock.Anything).Return(
		nil, errors.New("lorem")).Once()
	path := "/api/v1/green-skills"
	handler := rest.GreenSkillRESTHandler{Service: svcMock}
	app.Get(path, func(ctx *fiber.Ctx) error {
		ctx.Query("search", "lorem")
		ctx.Query("limit", "1")
		ctx.Query("page", "1")
		_ = ctx.Next()
		return nil
	}, handler.Fetch)
	req := httptest.NewRequest("GET", path, nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.StatusCode, fiber.StatusInternalServerError)
}

func Test_Skill_Show_ShouldSuccess(t *testing.T) {
	app := fiber.New()
	svcMock := new(mocks.ISkillService)
	svcMock.On("GreenSkillDetail",
		mock.Anything, mock.Anything,
	).Return(&response.GreenSkillResponse{ID: "123"}, nil).Once()
	path := "/api/v1/green-skills"
	handler := rest.GreenSkillRESTHandler{Service: svcMock}
	app.Get(fmt.Sprintf("%s/:id", path), func(ctx *fiber.Ctx) error {
		_ = ctx.Next()
		return nil
	}, handler.Show)
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/123-id", path), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.StatusCode, fiber.StatusOK)
}
func Test_Skill_Show_ShouldError(t *testing.T) {
	app := fiber.New()
	svcMock := new(mocks.ISkillService)
	svcMock.On("GreenSkillDetail",
		mock.Anything, mock.Anything,
	).Return(nil, errors.New("lorem")).Once()
	path := "/api/v1/green-skills"
	handler := rest.GreenSkillRESTHandler{Service: svcMock}
	app.Get(fmt.Sprintf("%s/:id", path), func(ctx *fiber.Ctx) error {
		_ = ctx.Next()
		return nil
	}, handler.Show)
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/123-id", path), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.StatusCode, fiber.StatusInternalServerError)
}
