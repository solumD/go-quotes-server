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

// config is a struct that contains the configuration for the server.
type config struct {
	loggerLevel string
	postgresDSN string
	serverHost  string
	serverPort  string
}

// MustLoad loads the config from the .env file.
func MustLoad() *config {
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

	return &config{
		loggerLevel: logLevel,
		postgresDSN: pgDSN,
		serverHost:  srvHost,
		serverPort:  srvPort,
	}
}

// LoggerLevel returns the logger level.
func (cfg *config) LoggerLevel() string {
	return cfg.loggerLevel
}

// PostgresDSN returns the DSN for the Postgres database connection.
func (cfg *config) PostgresDSN() string {
	return cfg.postgresDSN
}

// ServerAddr returns the server address by combining the server host and port.
func (cfg *config) ServerAddr() string {
	return net.JoinHostPort(cfg.serverHost, cfg.serverPort)
}
