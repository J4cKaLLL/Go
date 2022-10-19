package flag

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	// Inicialización con Singleton
	onceInit sync.Once
)

// Config  ...
type Config struct {
	GRPCPort      string
	HTTPPort      string
	LogLevel      int
	LogTimeFormat string
	Help          bool
}

// ShowHelp ...
func ShowHelp() {
	fmt.Println(`
Usage: CLI find_customer [OPTIONS]
Options:
	-g, --grpc-port		gRPC port para bind.		(default='50000').
	-p, --http-port		HTTP port para bind.		(default='8080').
	-l, --log-level		Global log level.		(default=0).
	-f, --log-format	Global format.			(default='2006-01-02T15:04:05Z07:00').
	-h, --help		Help.
    `)
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		ShowHelp()
		return
	}
}

const (
	defaultGrpcPort  = "50001"
	defaultHTTPPort  = "8080"
	defaultLogLevel  = "0"
	defaultLogFormat = "2006-01-02T15:04:05Z07:00"
)

// EnvValue ...
func EnvValue(env, defaultValue string) string {
	value := os.Getenv(env)
	if len(value) == 0 {
		return defaultValue
	}
	return value

}

// NewFlagConfig ...
func NewFlagConfig() *Config {

	config := &Config{}

	var (
		grpcPort    = EnvValue("GRPC_PORT", defaultGrpcPort)
		httpPort    = EnvValue("HTTP_PORT", defaultHTTPPort)
		logLevel, _ = strconv.Atoi(EnvValue("LOG_LEVEL", defaultLogLevel))
		logFormat   = EnvValue("LOG_FORMAT", defaultLogFormat)
	)

	onceInit.Do(func() {

		// gRPC flags
		flag.StringVar(&config.GRPCPort, "g", grpcPort, "gRPC port para bind")
		flag.StringVar(&config.GRPCPort, "grpc-port", grpcPort, "gRPC port para bind")

		// HTTP flags
		flag.StringVar(&config.HTTPPort, "p", httpPort, "HTTP port para bind")
		flag.StringVar(&config.HTTPPort, "http-port", httpPort, "HTTP port para bind")

		// Logs level flags
		flag.IntVar(&config.LogLevel, "l", logLevel, "Global log level")
		flag.IntVar(&config.LogLevel, "log-level", logLevel, "Global log level")

		// Logs format flags
		flag.StringVar(&config.LogTimeFormat, "f", logFormat, "Global format")
		flag.StringVar(&config.LogTimeFormat, "log-format", logFormat, "Global format")

		// Help flags
		flag.BoolVar(&config.Help, "h", false, "")
		flag.BoolVar(&config.Help, "help", false, "")

		setFlag(flag.CommandLine)

		flag.Parse()
	})

	return config

}
