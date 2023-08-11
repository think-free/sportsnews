package restylog

import (
	"context"
	"strings"

	"github.com/think-free/sportsnews/lib/logging"
)

// This logger is only used to log the internal errors of the resty client
// that can happen in the Set* functions.
// It is not used to log the request and response of the resty client.

type logWriter struct {
	ctx context.Context
}

func (e logWriter) Write(p []byte) (int, error) {
	// We try to log all the resty message that are not request or response
	// Request and response always starts with "---------------------" so we
	// are filtering strings that contains it.
	mes := string(p)
	if !strings.Contains(mes, "---------------------") {
		logging.L(e.ctx).Warnf("resty message : %s", mes)
	}
	return len(p), nil
}
