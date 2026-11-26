package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
}

type loader struct {
	errors []string
}

func (l *loader) integer(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		l.errors = append(l.errors, key+": phải là số")
		return def
	}
	return n
}

func (l *loader) text(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func (l *loader) required(key string, minLen int) string {
	v := os.Getenv(key)
	if len(v) < minLen {
		l.errors = append(l.errors, fmt.Sprintf("%s: bắt buộc, tối thiểu %d ký tự", key, minLen))
	}
	return v
}

func Load() (Config, error) {
	_ = godotenv.Load()
	loader := &loader{}
	if len(loader.errors) > 0 {
		return Config{}, fmt.Errorf("cấu hình sai:\n - %s", strings.Join(loader.errors, "\n - "))
	}
	port := loader.integer("PORT", 3000)
	return Config{Port: port}, nil
}
