package store

type Config struct {
	dburl string `toml:"db_url"`
}

func NewConfig() *Config {
	return &Config{}
}