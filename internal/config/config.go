package config

import (
	"log"
	"net"
	"os"
)

const (
	EnvLoggerLevel = "LOGGER_LEVEL"
	EnvPostgresDSN = "PG_DSN"
	EnvServerHost  = "SERVER_HOST"
	EnvServerPort  = "SERVER_PORT"
)

type Config struct {
	loggerLevel string
	postgresDSN string
	serverHost  string
	serverPort  string
}

func MustLoad() *Config {
	logLevel := os.Getenv(EnvLoggerLevel)
	if len(logLevel) == 0 {
		log.Fatal("logger level not found")
	}

	pgDSN := os.Getenv(EnvPostgresDSN)
	if len(pgDSN) == 0 {
		log.Fatal("postgres dsn not found")
	}

	srvHost := os.Getenv(EnvServerHost)
	if len(srvHost) == 0 {
		log.Fatal("server host not found")
	}

	srvPort := os.Getenv(EnvServerPort)
	if len(srvPort) == 0 {
		log.Fatal("server port not found")
	}

	return &Config{
		loggerLevel: logLevel,
		postgresDSN: pgDSN,
		serverHost:  srvHost,
		serverPort:  srvPort,
	}
}

func (cfg *Config) LoggerLevel() string {
	return cfg.loggerLevel
}

func (cfg *Config) PostgresDSN() string {
	return cfg.postgresDSN
}

func (cfg *Config) ServerAddr() string {
	return net.JoinHostPort(cfg.serverHost, cfg.serverPort)
}
