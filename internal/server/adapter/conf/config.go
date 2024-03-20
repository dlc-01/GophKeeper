package conf

import (
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"time"
)

type Config struct {
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
	Logger logger.ConfLogger
}
