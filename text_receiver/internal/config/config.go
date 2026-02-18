package config

import (
	"errors"
	"fmt"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/validation"
)

var ErrInvalidAppEnv = errors.New("app env is invalid or not supported")

type AppEnv string

const (
	Development AppEnv = "development"
	Release     AppEnv = "release"
)

func parseEnvVariable(env string) (AppEnv, error) {
	switch env {
	case string(Development):
		return Development, nil
	case string(Release):
		return Release, nil
	default:
		return "", ErrInvalidAppEnv
	}
}

type HTTPServerConfig struct {
	Port uint16
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type AppConfig struct {
	HTTPServer        *HTTPServerConfig
	Redis             *RedisConfig
	Env               AppEnv
	AnalyzerClientURL string
}

func New() (*AppConfig, error) {

	const op = "config.New"

	httpServerPort, err := getEnvAs[uint16]("HTTP_SERVER_PORT")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	envString, err := getEnv("APP_ENV")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	redisAddr, err := getEnv("REDIS_ADDR")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	analyzerClientURL, err := getEnv("ANALYZER_CLIENT_URL")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	redisPassword, err := getEnv("REDIS_PASSWORD")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	env, err := parseEnvVariable(envString)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	cfg := &AppConfig{
		HTTPServer: &HTTPServerConfig{
			Port: httpServerPort,
		},
		Redis: &RedisConfig{
			Addr:     redisAddr,
			Password: redisPassword,
		},
		Env:               env,
		AnalyzerClientURL: analyzerClientURL,
	}

	err = cfg.validate()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cfg, nil

}

func (c *AppConfig) validate() error {

	var validationErrors []error

	err := validation.Range(c.HTTPServer.Port, 1024, 65535)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "HTTPServer.Port", err))
	}

	err = validation.EmptyString(c.Redis.Password)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "Redis.Password", err))
	}

	_, err = validation.IsURI(c.Redis.Addr)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "Redis.Addr", err))
	}

	_, err = validation.IsURI(c.AnalyzerClientURL)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "AnalyzerClientURL", err))
	}

	return errors.Join(validationErrors...)
}
