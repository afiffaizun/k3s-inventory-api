package config

import (
	"os"
	"testing"
)

func TestGetEnv_WithEnvVar(t *testing.T) {
	os.Setenv("TEST_DB_HOST", "localhost")
	defer os.Unsetenv("TEST_DB_HOST")

	result := getEnv("TEST_DB_HOST", "default")
	if result != "localhost" {
		t.Errorf("expected 'localhost', got '%s'", result)
	}
}

func TestGetEnv_WithoutEnvVar(t *testing.T) {
	os.Unsetenv("TEST_DB_MISSING")
	result := getEnv("TEST_DB_MISSING", "fallback")
	if result != "fallback" {
		t.Errorf("expected 'fallback', got '%s'", result)
	}
}

func TestGetEnv_EmptyEnvVar(t *testing.T) {
	os.Setenv("TEST_DB_EMPTY", "")
	defer os.Unsetenv("TEST_DB_EMPTY")

	result := getEnv("TEST_DB_EMPTY", "fallback")
	if result != "" {
		t.Errorf("expected '', got '%s'", result)
	}
}
