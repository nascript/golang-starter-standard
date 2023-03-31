package account

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAccountProvider(_ fiber.Router, _ *mongo.Database) {}
