package wapp

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog"
)

// EnableProbes controls whether or not a call to Probes will start an HTTP server or not, since Probes can only be called once
var EnableProbes = true

// LivenessCheck defines something that is alive if it returns nil or not alive if it returns an error
type LivenessCheck interface {
	Liveness(ctx context.Context) error
}

type LivenessCheckFunc func(context.Context) error

func (r LivenessCheckFunc) Liveness(ctx context.Context) error {
	return r(ctx)
}

// Probes will create and run an HTTP server with endpoints for /liveness and /readiness. Any number of LivenessCheck instances can be passed in--if any returns an error then the app is not alive
func Probes(addr string, liveness ...LivenessCheck) StartFunc {
	// if EnableProbes is false then just return nil
	if EnableProbes == false {
		return nil
	}

	// set EnableProbes to false so that we only start one set of Probes
	EnableProbes = false

	return func() (StopFunc, error) {
		s := statusHandler{
			livenessChecks: liveness,
		}

		mux := http.NewServeMux()

		mux.HandleFunc("/liveness", s.Liveness)
		mux.HandleFunc("/readiness", s.Readiness)
		mux.HandleFunc("/appinfo", s.AppInfo)

		srv := &http.Server{Addr: addr, Handler: mux}
		go func() {
			_ = srv.ListenAndServe()
		}()

		return srv.Shutdown, nil
	}
}

type statusHandler struct {
	livenessChecks []LivenessCheck
}

// K8s hits this sometimes
// If liveness returns an error code K8S will restart the pod immediately
func (s statusHandler) Liveness(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	livenessChecks := append(s.livenessChecks, globalLivenessChecks...)

	for _, r := range livenessChecks {
		err := r.Liveness(ctx)
		if err != nil {
			wlog.WError(werror.Wrap(err, "error while checking liveness status"))
			http.Error(w, fmt.Sprintf("not alive: error: %s", err.Error()), http.StatusServiceUnavailable)
			return
		}
	}

	_, err := io.WriteString(w, "It's working")
	if err != nil {
		wlog.WError(werror.Wrap(err, "unable to write liveness body"))
		return
	}

}

// K8s hit on startup in k8s, once it returns a 20* code this won't get hit till the pod restarts
// If readiness returns an error code K8S will stop sending traffic to the pod
func (s statusHandler) Readiness(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "Ready to handle traffic")
	if err != nil {
		wlog.WError(werror.Wrap(err, "unable to write readiness body"))
		return
	}

}
