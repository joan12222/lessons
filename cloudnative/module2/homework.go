package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：
//
//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，应返回 200

func echoHeader(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	var vv string
	for k, v := range header {
		for _, vv = range v {
			w.Header().Add(k, vv)
		}
	}
	ver := os.Getenv("SystemDrive")
	w.Header().Set("VERSION", ver)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Request from %s, response code is %d", r.RemoteAddr, http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "sever is running"
	jsonresp, _ := json.Marshal(resp)
	w.Write(jsonresp)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Request from %s, response code is %d", r.RemoteAddr, http.StatusOK)
}

func main() {
	http.HandleFunc("/echoHeader", echoHeader)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}
