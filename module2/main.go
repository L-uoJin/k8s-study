package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"module2/config"
	"module2/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
func main() {
	root := http.NewServeMux()
	router.InitRouter(root)
	server := &http.Server{
		Addr:    config.Port,
		Handler: root,
	}
	go server.ListenAndServe()
	listenSignal(context.Background(), server)
	select {}
}
func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		log.Println("notify sigs")
		httpSrv.Shutdown(ctx)
		log.Println("http shutdown")
	}
}