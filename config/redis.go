package config

func (c *Config) InitRedisConn() {
	redisSingleton.Do(func() {
		// TODO IMPL
	})
}
