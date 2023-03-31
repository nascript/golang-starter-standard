package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"skilledin-green-skills-api/config"
	"skilledin-green-skills-api/pkg/http/wrapper"
	"skilledin-green-skills-api/pkg/token"
)

var (
	cookieToken string
	parseToken  *jwt.Token
	claimToken  *token.JSONWebTokenClaim
	err         error
	ok          bool
)

func Auth(ctx *fiber.Ctx) error {
	if cookieToken = ctx.Cookies("token"); cookieToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(wrapper.ResponseFormatWrapper{
			Error:   true,
			Code:    fiber.StatusUnauthorized,
			Message: "Invalid token",
			Data:    nil,
		})
	}

	if parseToken, err = jwt.ParseWithClaims(
		cookieToken,
		&token.JSONWebTokenClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Instance.AccessSecret), nil
		},
	); err != nil && !parseToken.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(wrapper.ResponseFormatWrapper{
			Error:   true,
			Code:    fiber.StatusUnauthorized,
			Message: "Token expired",
			Data:    err.Error(),
		})
	}

	if claimToken, ok = parseToken.Claims.(*token.JSONWebTokenClaim); !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(wrapper.ResponseFormatWrapper{
			Error:   true,
			Code:    fiber.StatusUnauthorized,
			Message: "Invalid token",
			Data:    err.Error(),
		})
	}

	var userID string
	if uid, ok := claimToken.Payload.(map[string]any)["user_id"].(string); ok {
		userID = uid
	}

	ctx.Locals("user_id", userID)
	return ctx.Next()
}
