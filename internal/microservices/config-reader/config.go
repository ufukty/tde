package config_reader

import "time"

type Common struct {
	ServiceDiscoveryConfig       string        `yaml:"service-discovery-config"`
	ServiceDiscoveryUpdatePeriod time.Duration `yaml:"service-discovery-update-period"`
	GracePeriod                  time.Duration `yaml:"grace-period"`
	RouterPublic                 string        `yaml:"router_public"`
	RouterPrivate                string        `yaml:"router_private"`
}

type APIGateway struct {
	Common `yaml:",inline"`
}

type Captcha struct {
	Common `yaml:",inline"`
}

type Customs struct {
	Common    `yaml:",inline"`
	MountPath string `yaml:"volume"`
}

type Evolver struct {
	Common `yaml:",inline"`
}

type Runner struct {
	Common `yaml:",inline"`
}

type Config struct {
	Common     Common     `yaml:"common"`
	APIGateway APIGateway `yaml:"api_gateway"`
	Captcha    Captcha    `yaml:"captcha"`
	Customs    Customs    `yaml:"customs"`
	Evolver    Evolver    `yaml:"evolver"`
	Runner     Runner     `yaml:"runner"`
}
