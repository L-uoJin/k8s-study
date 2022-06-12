package util

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息
func  WithMsg(message string,code int,data map[string]interface{})  []byte {
	resp:= Response{
		Code: code,
		Msg:  message,
		Data: data,
	}
	respByte,err:=json.Marshal(&resp)
	if err!=nil{
		log.Error("json fail",err)
	}
	return  respByte
}
