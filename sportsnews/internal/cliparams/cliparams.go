package cliparams

import (
	"github.com/kelseyhightower/envconfig"
)

type ClientParameters struct {
	Environment   string `envconfig:"ENVIRONMENT" required:"false" default:"production"`
	LogLevel      string `envconfig:"LOGLEVEL" required:"false" default:"debug"`
	DatabaseURL   string `envconfig:"DATABASE_URL" required:"false" default:"mongodb://mongodb:27017"`
	ListenAddress string `envconfig:"LISTEN_ADDRESS" required:"false" default:"0.0.0.0:8080"`
}

func New() *ClientParameters {
	cp := &ClientParameters{}
	envconfig.MustProcess("SPORTSNEWS", cp)

	return cp
}
