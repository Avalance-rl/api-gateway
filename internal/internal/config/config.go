package config

import (
	"fmt"

	baseConfig "github.com/avalance-rl/otiva-pkg/config"
	"github.com/spf13/viper"
)

type ApiGatewayConfig struct {
	*baseConfig.Config
	AuthService struct {
		Host string
		Port string
	}
}

func Load(path string) (*ApiGatewayConfig, error) {
	baseCFG, err := baseConfig.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load base config: %w", err)
	}

	config := &ApiGatewayConfig{
		Config: baseCFG,
	}

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read auth_service config: %w", err)
	}

	if err := viper.UnmarshalKey("auth_service", &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal auth_service config: %w", err)
	}
	return config, nil
}
