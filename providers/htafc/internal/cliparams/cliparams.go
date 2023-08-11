package cliparams

import (
	"github.com/kelseyhightower/envconfig"
)

type ClientParameters struct {
	Environment string `envconfig:"ENVIRONMENT" required:"false" default:"production"`
	LogLevel    string `envconfig:"LOGLEVEL" required:"false" default:"debug"`
	DatabaseURL string `envconfig:"DATABASE_URL" required:"false" default:"mongodb://mongodb:27017"`

	MockedDatabase bool `envconfig:"MOCKED_DATABASE" required:"false" default:"true"` // Only used when environment is 'mock'
	MockedUpstream bool `envconfig:"MOCKED_UPSTREAM" required:"false" default:"true"` // Only used when environment is 'mock'
}

func New() *ClientParameters {
	cp := &ClientParameters{}
	envconfig.MustProcess("SPORTSNEWS", cp)

	return cp
}
