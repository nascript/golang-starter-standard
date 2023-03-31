package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"skilledin-green-skills-api/config"
	"skilledin-green-skills-api/modules"

	// docs
	_ "skilledin-green-skills-api/docs"
)

// @title Skilledin API Documentation
// @version 1.0
// @description This is an auto-generated API Docs. To Using Private API 🔐 just hit endpoint AUTH > Login. We Are using cookie header no need copy token.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email developer@skilledin.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	// SET .env File
	viper.SetConfigFile(".env")
	// LOAD CONFIG
	config.LoadEnv()
	// INIT DATABASE CONNECTION
	config.Instance.InitMongoDBConn()
	// INIT REDIS CONNECTION
	config.Instance.InitRedisConn()
	// INIT FIBER FRAMEWORK WITH BASIC MIDDLEWARE
	config.Instance.InitFiber()
}

func main() {
	// REGISTER MODULES
	modules.Register(
		modules.WithEngine(config.FiberApp),
		modules.WithDatabase(config.DbPool))
	// RUN THE APP
	log.Fatalln(config.FiberApp.Listen(fmt.Sprintf(
		":%s", config.Instance.DevServerURL)))
}
