package skill

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"skilledin-green-skills-api/modules/skill/common"
	"skilledin-green-skills-api/modules/skill/delivery/rest"
	mongoRepository "skilledin-green-skills-api/modules/skill/repository/mongo"
	"skilledin-green-skills-api/modules/skill/service"
)

func NewSkillProvider(router fiber.Router, database *mongo.Database) {
	skillCollection := database.Collection(common.SkillCollectionName)
	greenSkillCollection := database.Collection(common.GreenSkillCollectionName)
	greenSkillGroupCollection := database.Collection(common.GreenSkillGroupCollectionName)
	skillRepository := mongoRepository.NewSkillMongoRepository(
		mongoRepository.WithSkillCollection(skillCollection),
		mongoRepository.WithGreenSkillCollection(greenSkillCollection),
		mongoRepository.WithGreenSkillGroupCollection(greenSkillGroupCollection))
	skillService := service.NewSkillService(skillRepository)
	// TODO ADD AUTH MIDDLEWARE AND PROTECT THE ROUTES
	greenSkillRouterGroup := router.Group("/green-skills")
	rest.NewGreenSkillRESTHandler(greenSkillRouterGroup, skillService)
	skillRouterGroup := router.Group("/skills")
	rest.NewSkillRESTHandler(skillRouterGroup, skillService)
}
