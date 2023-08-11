package restylog

import (
	"net/http"

	"gopkg.in/resty.v1"
)

type LogFilterFct func(req *resty.Request, resp *resty.Response) bool

func LogStatusNonOk(req *resty.Request, resp *resty.Response) bool {
	return resp.StatusCode() != http.StatusOK
}

func LogEverything(req *resty.Request, resp *resty.Response) bool {
	return true
}

func LogNothing(req *resty.Request, resp *resty.Response) bool {
	return false
}
