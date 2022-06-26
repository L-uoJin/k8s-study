package controller

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"module2/controller/metrics"
	"module2/util"
	"net"
	"net/http"
	"os"
	"time"
)

func Healthz(w http.ResponseWriter,r *http.Request)  {
	for k,v:=range r.Header{
		w.Header().Set(fmt.Sprintf("%v",k),v[0])
		log.Info("header key is ",k," value is ",v)
	}
	version:= os.Getenv("VERSION")
	log.Info("VERSION is ",version)
	w.Header().Set("VERSION",version)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Error("Client IP get err:", err)
	}else{
		log.Info("Client IP is ",ip)
	}
	log.Info("http code is ",http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write(util.WithMsg("sucess",http.StatusOK,
		map[string]interface{}{
		"version":version,
		"clientIp":ip}))
}

func Rand(w http.ResponseWriter,r *http.Request)  {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond*time.Duration(delay))
	w.WriteHeader(http.StatusOK)
	w.Write(util.WithMsg("sucess",http.StatusOK, nil))
}
func randInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
