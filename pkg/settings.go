package pkg

type Config struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_DBNAME"`
	DBSSLMode  string `env:"DB_SSLMODE"`
	Port       string `env:"APP_PORT"`
}
