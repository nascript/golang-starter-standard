package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"skilledin-green-skills-api/modules/account"
	"skilledin-green-skills-api/modules/job"
	"skilledin-green-skills-api/modules/skill"
)

type Module struct {
	Engine   *fiber.App
	Database *mongo.Database
}

type ModuleOption func(*Module)

func WithEngine(engine *fiber.App) ModuleOption {
	return func(module *Module) {
		module.Engine = engine
	}
}

func WithDatabase(database *mongo.Database) ModuleOption {
	return func(module *Module) {
		module.Database = database
	}
}

func Register(options ...ModuleOption) {
	module := &Module{}
	for _, option := range options {
		option(module)
	}
	module.registerPublicRouteHandler()
	module.registerModules()
}

func (m *Module) registerPublicRouteHandler() {
	router := m.Engine
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Hello World!",
		})
	})
	router.Get("/swagger/*", swagger.HandlerDefault)
}

func (m *Module) registerModules() {
	router := m.Engine.Group("api/v1")
	database := m.Database

	account.NewAccountProvider(router, database)
	skill.NewSkillProvider(router, database)
	job.NewJobProvider(router, database)
}
