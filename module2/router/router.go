package router

import (
	"module2/controller"
	"net/http"
)

func InitRouter(mux *http.ServeMux)  {
	mux.HandleFunc("/healthz", controller.Healthz)
	mux.HandleFunc("/rand", controller.Rand)
	mux.HandleFunc("/healthz", controller.Healthz)
}
