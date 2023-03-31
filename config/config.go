package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

var (
	cfgSingleton   sync.Once
	fiberSingleton sync.Once
	mongoSingleton sync.Once
	redisSingleton sync.Once

	Instance *Config
	FiberApp *fiber.App
	DbPool   *mongo.Database
	// RedisPool *redis
)

type Config struct {
	DevServerURL           string `mapstructure:"DEV_SERVER_URL"`
	MongoURI               string `mapstructure:"MONGOURI"`
	MongoDB                string `mapstructure:"MONGO_DATABASE"`
	MongoPoolMin           int    `mapstructure:"MONGO_POOL_MIN"`
	MongoPoolMax           int    `mapstructure:"MONGO_POOL_MAX"`
	MongoMaxIdleTimeSecond int    `mapstructure:"MONGO_MAX_IDLE_TIME_SECOND"`
	AccessSecret           string `mapstructure:"ACCESS_SECRET"`
	SendgridAPIKey         string `mapstructure:"SENDGRID_API_KEY"`
	EmailDev               string `mapstructure:"EMAIL_DEVELOPER_SKILLEDIN"`
	StripeKey              string `mapstructure:"STRIPE_KEY_TEST"`
	RedirectVerifyEmail    string `mapstructure:"URL_REDIRECT_VERIFY_EMAIL"`
	RedirectInviteEmail    string `mapstructure:"URL_REDIRECT_INVITE_EMAIL"`
}

func LoadEnv() {
	// notify that app try to load config file
	log.Println("Load configuration file . . . .")

	cfgSingleton.Do(func() {
		// find environment file
		viper.AutomaticEnv()
		// error handling for specific case
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				panic(".env file not found!, please copy .env.example and paste as .env")
			}

			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
		// notify that config file is ready
		log.Println("configuration file: ready")
		// extract config to struct
		if err := viper.Unmarshal(&Instance); err != nil {
			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
	})
}
