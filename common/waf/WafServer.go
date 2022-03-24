package waf

import (
	"context"
	"net/http"
	"time"
)

// waf server
// author wuxw
type WafServer struct {
	server http.Server
}

// start waf
func (waf *WafServer) StartWaf(port string)  error{

	//root handle
	http.HandleFunc("/", RootHandle)

	waf.server = http.Server{
		Addr: ":"+port,
		Handler: http.DefaultServeMux,
	}
	err := waf.server.ListenAndServe()
	return err
}

// stop waf
func (waf *WafServer) StopWaf() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := waf.server.Shutdown(ctx)
	return err
}
