/*
  whttpclient does http requests and injects tracing and other meta data into the outbound request.
  We need to decide if this is a one off thing for this service, our if it should be added to
  wlib.
*/
package whttpclient

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
	"weavelab.xyz/monorail/shared/wlib/wtracer"
)

var client = http.DefaultClient

const clientMetricName = "whttpclient"

func init() {
	client.Transport = &nethttp.Transport{}

	wmetrics.SetLabels(clientMetricName, "Host", "Method", "StatusCode")
}

func Get(ctx context.Context, u string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	r, err := Do(ctx, req)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return r, nil

}

func Post(ctx context.Context, u string, contentType string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	r, err := Do(ctx, req)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	return r, nil

}

func PostForm(ctx context.Context, url string, data url.Values) (*http.Response, error) {
	return Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func Do(ctx context.Context, req *http.Request) (*http.Response, error) {

	host := strings.Replace(req.URL.Host, ".", "_", -1)

	statusCode := "unknown"
	timer := wmetrics.StartTimer(clientMetricName, host, req.Method)
	defer func() {
		timer(statusCode)
	}()

	// add context to the request
	req = req.WithContext(ctx)

	// add tracing headers if present
	tracer, err := wtracer.DefaultTracer()
	if err != nil {
		return nil, werror.Wrap(err)
	}

	req, ht := nethttp.TraceRequest(tracer, req)
	defer ht.Finish()

	resp, err := client.Do(req)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)

	return resp, nil

}
