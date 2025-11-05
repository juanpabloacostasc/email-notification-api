package config_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"notification-service/config"
	"notification-service/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	validConfig := `{"rate_limit_rules": [{"type": "news", "limit": 1, "duration": "1h"}]}`
	invalidConfig := `{"rate_limit_rules": [`

	tempDir := t.TempDir()

	validConfigFile := filepath.Join(tempDir, "valid_config.json")
	err := os.WriteFile(validConfigFile, []byte(validConfig), 0644)
	require.NoError(t, err)

	invalidConfigFile := filepath.Join(tempDir, "invalid_config.json")
	err = os.WriteFile(invalidConfigFile, []byte(invalidConfig), 0644)
	require.NoError(t, err)

	tests := []struct {
		name        string
		path        string
		expectError bool
		expected    *config.Config
	}{
		{
			name:        "success - valid config",
			path:        validConfigFile,
			expectError: false,
			expected: &config.Config{
				RateLimitRules: []config.RateLimitRule{
					{Type: "news", Limit: 1, Duration: "1h"},
				},
			},
		},
		{
			name:        "fail - file not found",
			path:        "non_existent_file.json",
			expectError: true,
			expected:    nil,
		},
		{
			name:        "fail - invalid json",
			path:        invalidConfigFile,
			expectError: true,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := config.LoadConfig(tt.path)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, cfg)
			}
		})
	}
}

func TestToDomainRateLimitRules(t *testing.T) {
	tests := []struct {
		name     string
		config   *config.Config
		expected []domain.RateLimitRule
	}{
		{
			name: "success - valid config",
			config: &config.Config{
				RateLimitRules: []config.RateLimitRule{
					{Type: "news", Limit: 1, Duration: "1h"},
					{Type: "status", Limit: 2, Duration: "1m"},
				},
			},
			expected: []domain.RateLimitRule{
				{Type: "news", Limit: 1, Duration: time.Hour},
				{Type: "status", Limit: 2, Duration: time.Minute},
			},
		},
		{
			name: "success - invalid duration",
			config: &config.Config{
				RateLimitRules: []config.RateLimitRule{
					{Type: "news", Limit: 1, Duration: "invalid"},
				},
			},
			expected: []domain.RateLimitRule{
				{Type: "news", Limit: 1, Duration: 0},
			},
		},
		{
			name:     "success - empty config",
			config:   &config.Config{},
			expected: []domain.RateLimitRule{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := tt.config.ToDomainRateLimitRules()
			assert.Equal(t, tt.expected, rules)
		})
	}
}