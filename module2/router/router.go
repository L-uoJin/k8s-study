package router

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"module2/controller"
	"net/http"
)

func InitRouter(mux *http.ServeMux)  {
	mux.HandleFunc("/healthz", controller.Healthz)
	mux.HandleFunc("/rand", controller.Rand)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", controller.Healthz)
}
