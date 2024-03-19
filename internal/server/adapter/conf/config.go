package conf

import "time"

type Config struct {
	AppMod          string        `env:"APP_ENV" envDefault:"develop" json:"AppMod"`
	Config          string        `env:"CONF" envDefault:""`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"600s" json:"ShutdownTimeout"`
	GRPCServer      struct {
		Address string `env:"GRPC_SERVER_ADDRESS" envDefault:":3200" json:"GRPCAddress"`
		Network string `env:"GRPC_SERVER_NETWORK" envDefault:"tcp" json:"GRPCNetwork"`
	}
	DB struct {
		DSN    string `env:"DATABASE_DSN" envDefault:"postgresql://admin:root@localhost:5432/postgres" json:"DB_DSN"`
		Driver string `env:"DATABASE_DRIVER" envDefault:"pgx" json:"DB_Driver"`
	}
	JWT struct {
		Expire    time.Duration `env:"JWT_TOKEN_EXPIRE" envDefault:"10m" json:"JWTTokenExpire"`
		SecretKey string        `env:"JWT_SECRET_KEY" envDefault:"lol" json:"JwtSecretKey"`
	}
	Logger struct {
		File struct {
			Directory  string `env:"LOGGER_DIRECTORY" envDefault:"temp/logs/" json:"LoggerDirectory"`
			MaxSize    int    `env:"LOGGER_FILE_MAX_SIZE" envDefault:"1" json:"LoggerMaxSize"`
			MaxBackups int    `env:"LOGGER_FILE_MAX_BACKUPS" envDefault:"1" json:"LoggerMaxBackups"`
			MaxAge     int    `env:"LOGGER_FILE_MAX_AGE" envDefault:"1" json:"LoggerMaxAge"`
			Compress   bool   `env:"LOGGER_FILE_COMPRESS" envDefault:"true" json:"LoggerCompress"`
		}
	}
}
