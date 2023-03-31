package rest

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"skilledin-green-skills-api/modules/skill/common"
	"skilledin-green-skills-api/modules/skill/domain"
	"skilledin-green-skills-api/pkg/http/parser"
	"skilledin-green-skills-api/pkg/http/wrapper"
	"time"
)

type GreenSkillRESTHandler struct {
	Service domain.ISkillService
}

// Fetch
// Green Skill godoc
// @Schemes
// @Summary		green skill list
// @Description	get green skill data list
// @Tags		Green Skills
// @Accept		json
// @Produce		json
// @Param 		limit  			query    int  		false "Limit default: 10"
// @Param 		page   			query    int  		false "Page default: 1"
// @Param 		search   		query    string 	false "Search data by title/name"
// @Param 		groups   		query    []string  	false "Filter by skill groups name"
// @Param 		sort_rate 		query    string  	false "Sort order by ratings default:none" Enums(asc, desc)
// @Param 		sort_trf 		query    string  	false "Sort order by transferability default:none, asc for low to high, and desc for high to low" Enums(asc, desc)
// @Security 	ApiKeyAuth
// @Router /api/v1/green-skills [GET]
func (handler *GreenSkillRESTHandler) Fetch(ctx *fiber.Ctx) error {
	limit, offset := parser.ParsePaginationQueryParam(ctx)
	search := ctx.Query("search")
	ctxWT, cancel := context.WithTimeout(ctx.Context(), time.Second*common.ContextTimeout)
	defer cancel()
	data, err := handler.Service.GreenSkillList(ctxWT, limit, offset, search)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(wrapper.ResponseFormatWrapper{
			Error:   true,
			Code:    fiber.StatusInternalServerError,
			Message: http.StatusText(fiber.StatusInternalServerError),
			Data:    err.Error(),
		})
	}

	data.ParseData(ctx)
	return ctx.Status(fiber.StatusOK).JSON(wrapper.ResponseFormatWrapper{
		Error:   false,
		Code:    fiber.StatusOK,
		Message: http.StatusText(fiber.StatusOK),
		Data:    data,
	})
}

// Show
// Green Skill godoc
// @Schemes
// @Summary		green skill detail
// @Description	get green skill detail
// @Tags		Green Skills
// @Accept		json
// @Produce		json
// @Param 		id			path	string  true "green_skill_id"
// @Security 	ApiKeyAuth
// @Router /api/v1/green-skills/{id} [GET]
func (handler *GreenSkillRESTHandler) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	ctxWT, cancel := context.WithTimeout(ctx.Context(), time.Second*common.ContextTimeout)
	defer cancel()
	data, err := handler.Service.GreenSkillDetail(ctxWT, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(wrapper.ResponseFormatWrapper{
			Error:   true,
			Code:    fiber.StatusInternalServerError,
			Message: http.StatusText(fiber.StatusInternalServerError),
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(wrapper.ResponseFormatWrapper{
		Error:   false,
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

func NewGreenSkillRESTHandler(
	router fiber.Router,
	service domain.ISkillService,
) {
	handler := &GreenSkillRESTHandler{service}
	router.Get("/", handler.Fetch)
	router.Get("/:id", handler.Show)
}
