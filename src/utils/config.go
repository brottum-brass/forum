package utils

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port            int
	DefaultLanguage string
	DefaultTheme    string
	CSS             string
	Icons           string
	JS              string
	Locales         string
	Themes          string
}

func NewConfig() *Config {
	port := mustGetEnv("PORT")
	defaultLanguage := mustGetEnv("DEFAULT_LANGUAGE")
	defaultTheme := mustGetEnv("DEFAULT_THEME")
	css := mustGetEnv("CSS")
	icons := mustGetEnv("ICONS")
	js := mustGetEnv("JS")
	locales := mustGetEnv("LOCALES")
	themes := mustGetEnv("THEMES")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Error converting PORT environment variable to integer: %v", err)
	}

	return &Config{
		Port:            portInt,
		DefaultLanguage: defaultLanguage,
		DefaultTheme:    defaultTheme,
		CSS:             css,
		Icons:           icons,
		JS:              js,
		Locales:         locales,
		Themes:          themes,
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}

	return value
}
