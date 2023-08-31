package config

import (
	"github.com/team-four-fingers/kakao/local"
	"github.com/team-four-fingers/kakao/mobility"
)

type Config struct {
	setting     *Setting
	mobilityCli mobility.Client
	localCli    local.Client
}

func (c *Config) LocalCli() local.Client {
	return c.localCli
}

func (c *Config) MobilityCli() mobility.Client {
	return c.mobilityCli
}

func (c *Config) Setting() *Setting {
	return c.setting
}

func NewConfig(setting *Setting, mobilityCli mobility.Client, localCli local.Client) *Config {
	return &Config{
		setting:     setting,
		mobilityCli: mobilityCli,
		localCli:    localCli,
	}
}
