package job

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewJobProvider(_ fiber.Router, _ *mongo.Database) {}
