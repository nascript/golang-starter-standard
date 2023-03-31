package middleware_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"skilledin-green-skills-api/config"
	"skilledin-green-skills-api/pkg/http/middleware"
	"skilledin-green-skills-api/pkg/token"
	"testing"
	"time"
)

func TestAuthMiddleware(t *testing.T) {
	viper.Reset()
	viper.SetConfigFile("../../../.env")
	viper.SetConfigType("dotenv")
	config.LoadEnv()

	app := fiber.New()
	app.Use(middleware.Auth)

	t.Run("ERROR COOKIE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res, err := app.Test(req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("ERROR PARSE TOKEN WITH CLAIM", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJQT1NCRV9URVNUIiwiZXhwIjotNjIxMzU1OTY4MDAsImlhdCI6LTYyMTM1NTk2ODAwLCJwYXlsb2FkIjp7ImRhdGEiOiJoZWxsbyB3b3JsZCJ9fQ.-_tfeKKhqSRP2H_pVg4f_spkX_Z1Lo1nuiu09OFFvO0"
		req.AddCookie(&http.Cookie{Name: "token", Value: accessToken})
		res, err := app.Test(req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("SUCCESS", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		jwt := token.JSONWebToken{
			Issuer:    "skilledin.io",
			SecretKey: []byte(config.Instance.AccessSecret),
			IssuedAt:  time.Now(),
			ExpiredAt: time.Now().Add(1 * time.Minute),
		}
		accessToken, err := jwt.Claim(map[string]any{
			"user_id": "123-456",
		})
		assert.Nil(t, err)
		req.AddCookie(&http.Cookie{Name: "token", Value: accessToken})
		res, err := app.Test(req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
