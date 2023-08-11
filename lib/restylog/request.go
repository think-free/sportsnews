package restylog

import (
	"context"
	"fmt"
	"time"

	"github.com/think-free/sportsnews/lib/logging"
	"gopkg.in/resty.v1"
)

type Request struct {
	request   *resty.Request
	tag       string
	logFilter LogFilterFct
}

func CreateRequest(ctx context.Context, r *resty.Request, tag string) *Request {
	r.SetContext(ctx)
	return &Request{
		r,
		tag,
		LogStatusNonOk,
	}
}

func (r *Request) SetContext(ctx context.Context) *Request {
	r.request.SetContext(ctx)
	return r
}

func (r *Request) Get(url string) (*Response, error) {
	res, err := r.request.Get(url)
	r.log("GET", url, res)
	return &Response{res}, err
}

func (r *Request) Head(url string) (*Response, error) {
	res, err := r.request.Head(url)
	r.log("HEAD", url, res)
	return &Response{res}, err
}

func (r *Request) Post(url string) (*Response, error) {
	res, err := r.request.Post(url)
	r.log("POST", url, res)
	return &Response{res}, err
}

func (r *Request) Put(url string) (*Response, error) {
	res, err := r.request.Put(url)
	r.log("PUT", url, res)
	return &Response{res}, err
}

func (r *Request) Delete(url string) (*Response, error) {
	res, err := r.request.Delete(url)
	r.log("DELETE", url, res)
	return &Response{res}, err
}

func (r *Request) Patch(url string) (*Response, error) {
	res, err := r.request.Patch(url)
	r.log("PATCH", url, res)
	return &Response{res}, err
}

func (r *Request) Options(url string) (*Response, error) {
	res, err := r.request.Options(url)
	r.log("OPTION", url, res)
	return &Response{res}, err
}

// SetLogFilter sets a custom filter for the logging
// The function passed as argument must return true if the request and response must be logged
func (r *Request) SetLogFilter(f func(req *resty.Request, resp *resty.Response) bool) *Request {
	r.logFilter = f
	return r
}

func (r *Request) log(method, url string, res *resty.Response) {
	ctx := r.request.Context()
	ts := time.Now().UnixNano()
	// We apply the filter only on the header and body of the request/response
	if r.logFilter(r.request, res) {
		// Request
		for k, v := range r.request.Header {
			ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.request.header.%s", r.tag, ts, k), v)
		}
		ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.request.body", r.tag, ts), r.request.Body)
		ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.request.queryparams", r.tag, ts), r.request.QueryParam)
		// Response
		if res != nil {
			for k, v := range res.Header() {
				ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.response.header.%s", r.tag, ts, k), v)
			}
			ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.response.body", r.tag, ts), res.String())
		}
	}

	ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.request.method", r.tag, ts), method)
	ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.request.url_path", r.tag, ts), url)
	if res != nil {
		ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.request.url", r.tag, ts), res.Request.URL)
		ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.response.status", r.tag, ts), res.StatusCode())
		ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.response.time", r.tag, ts), res.ReceivedAt().Format(time.RFC3339Nano))
		ctx = logging.SetTag(ctx, fmt.Sprintf("resty.%s.%d.response.duration", r.tag, ts), res.Time())
	}
	logging.L(ctx).Debugf("%s : %s", method, url)
}

func (r *Request) SetHeader(k, v string) *Request {
	r.request.SetHeader(k, v)
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.request.SetBody(body)
	return r
}

func (r *Request) SetError(err interface{}) *Request {
	r.request.SetError(err)
	return r
}

func (r *Request) SetResult(result interface{}) *Request {
	r.request.SetResult(result)
	return r
}

func (r *Request) SetBasicAuth(username, password string) *Request {
	r.request.SetBasicAuth(username, password)
	return r
}

func (r *Request) SetQueryParam(k, v string) *Request {
	r.request.SetQueryParam(k, v)
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	r.request.SetQueryParams(params)
	return r
}

func (r *Request) SetPathParams(params map[string]string) *Request {
	r.request.SetPathParams(params)
	return r
}

func (r *Request) SetFormData(data map[string]string) *Request {
	r.request.SetFormData(data)
	return r
}
