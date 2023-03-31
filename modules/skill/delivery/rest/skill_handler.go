package rest

import (
	"github.com/gofiber/fiber/v2"
	"skilledin-green-skills-api/modules/skill/domain"
)

type SkillRESTHandler struct {
	Service domain.ISkillService
}

func (handler *SkillRESTHandler) Fetch(_ *fiber.Ctx) error {
	return nil
}

func (handler *SkillRESTHandler) Create(_ *fiber.Ctx) error {
	return nil
}

func NewSkillRESTHandler(
	router fiber.Router,
	service domain.ISkillService,
) {
	handler := &SkillRESTHandler{service}
	router.Get("/", handler.Fetch)
	router.Post("/", handler.Create)
}
