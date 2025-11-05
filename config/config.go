package config

import (
	"encoding/json"
	"os"
	"time"

	"notification-service/internal/domain"
)

type Config struct {
	RateLimitRules []RateLimitRule `json:"rate_limit_rules"`
}

type RateLimitRule struct {
	Type     string `json:"type"`
	Limit    int    `json:"limit"`
	Duration string `json:"duration"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) ToDomainRateLimitRules() []domain.RateLimitRule {
	rules := make([]domain.RateLimitRule, len(c.RateLimitRules))
	for i, r := range c.RateLimitRules {
		duration, _ := time.ParseDuration(r.Duration)
		rules[i] = domain.RateLimitRule{
			Type:     r.Type,
			Limit:    r.Limit,
			Duration: duration,
		}
	}
	return rules
}