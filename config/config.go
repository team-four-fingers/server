package config

import (
	"github.com/team-four-fingers/kakao/mobility"
)

type Config struct {
	setting     *Setting
	mobilityCli mobility.Client
}

func (c *Config) MobilityCli() mobility.Client {
	return c.mobilityCli
}

func (c *Config) Setting() *Setting {
	return c.setting
}

func NewConfig(setting *Setting, mobilityCli mobility.Client) *Config {
	return &Config{
		setting:     setting,
		mobilityCli: mobilityCli,
	}
}
