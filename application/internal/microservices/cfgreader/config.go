package cfgreader

import "time"

type RouterParameters struct {
	RequestTimeout time.Duration `yaml:"request-timeout"`
	WriteTimeout   time.Duration `yaml:"write-timeout"`
	ReadTimeout    time.Duration `yaml:"read-timeout"`
	IdleTimeout    time.Duration `yaml:"idle-timeout"`
	GracePeriod    time.Duration `yaml:"grace-period"`
}

type Common struct {
	RouterParameters             `yaml:",inline"`
	ServiceDiscoveryConfig       string        `yaml:"service-discovery-config"`
	ServiceDiscoveryUpdatePeriod time.Duration `yaml:"service-discovery-update-period"`
	RouterPublic                 string        `yaml:"router-public"`
	RouterPrivate                string        `yaml:"router-private"`
}

type APIGateway struct {
	Common `yaml:",inline"`
}

type Captcha struct {
	Common `yaml:",inline"`
}

type Customs struct {
	Common    `yaml:",inline"`
	MountPath string `yaml:"volume-mount"`
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
