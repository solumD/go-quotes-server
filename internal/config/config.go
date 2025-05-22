package config

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

const (
	configPath     = ".env"
	envLoggerLevel = "LOGGER_LEVEL"
	envPostgresDSN = "PG_DSN"
	envServerHost  = "SERVER_HOST"
	envServerPort  = "SERVER_PORT"
)

type Config struct {
	loggerLevel string
	postgresDSN string
	serverHost  string
	serverPort  string
}

func MustLoad() *Config {
	err := godotenv.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load %s: %v", configPath, err)
	}

	logLevel := os.Getenv(envLoggerLevel)
	if len(logLevel) == 0 {
		log.Fatal("logger level not found")
	}

	pgDSN := os.Getenv(envPostgresDSN)
	if len(pgDSN) == 0 {
		log.Fatal("postgres dsn not found")
	}

	srvHost := os.Getenv(envServerHost)
	if len(srvHost) == 0 {
		log.Fatal("server host not found")
	}

	srvPort := os.Getenv(envServerPort)
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
