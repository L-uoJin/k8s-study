package config

import (
	"fmt"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)
var (
	Port  string
	LogLevel string
)
func init()  {
	cfgs, err := ini.Load("/module2/conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	Port = cfgs.Section("web").Key("port").Value()
	LogLevel = cfgs.Section("web").Key("loglevel").Value()
	level, err := log.ParseLevel(LogLevel)
	if err != nil {
		level = log.GetLevel()
	}
	log.SetLevel(level)
}
