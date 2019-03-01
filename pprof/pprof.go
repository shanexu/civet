package pprof

import (
	"net/http"
	"net/http/pprof"
)

func ListenAndServe(addr string, prefix ...string) error {
	var h string
	if len(prefix) == 0 {
		h = "/debug"
	} else {
		h = prefix[0]
	}
	h = h + "/pprof"
	mux := http.NewServeMux()
	mux.HandleFunc(h, pprof.Index)
	mux.HandleFunc(h +"/cmdline", pprof.Cmdline)
	mux.HandleFunc(h + "/profile", pprof.Profile)
	mux.HandleFunc(h + "/symbol", pprof.Symbol)
	mux.HandleFunc(h + "/trace", pprof.Trace)
	return http.ListenAndServe(addr, mux)
}
