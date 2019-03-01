package pprof

import (
	"net/http"
	"testing"
	"time"
)

func TestListenAndServe(t *testing.T) {
	go func() {
		ListenAndServe(":8089", "/test")
	} ()
	time.Sleep(time.Millisecond * 100)
	resp, err := http.DefaultClient.Get("http://127.0.0.1:8089/test/pprof/symbol")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("status code is not 200, but %d", resp.StatusCode)
	}
}
