package endpoint

import (
	"net/http"
	"crypto/tls"
	"log"
	"net"
	"time"
	"io"
	"io/ioutil"
	"strconv"
	"apibuilder-server/model"
	"encoding/json"
)
var proxyChannelId int
func Proxy(port, channelId int) {
	proxyChannelId = channelId
	addr := ":"+strconv.Itoa(port)
	log.Print("Test Listening and Proxy serving on " + addr)
	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				handleTunneling(w, r)
			} else {
				handleHTTP(w, r)
			}
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	server.ListenAndServe()
}
func handleTunneling(w http.ResponseWriter, r *http.Request) {
	by,_ := ioutil.ReadAll(r.Body)
	log.Print(r.RequestURI, string(by),r.Header)
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
func handleHTTP(w http.ResponseWriter, req *http.Request) {
	var p []byte
	t,_ := req.Body.Read(p)
	log.Print(req.RequestURI, t, string(p), req.Header["Content-Type"][0], req.Header["Authorization"][0])
	//to create model
	proxyRequest := &model.ProxyReq{}
	proxyRequest.ProxyChannelId = proxyChannelId
	proxyRequest.RemoteAddr = req.RemoteAddr
	proxyRequest.UserAgent = req.UserAgent()
	proxyRequest.RequestUrl = req.RequestURI
	proxyRequest.Method = req.Method
	hader :=  model.JSON{}
	hadbyte,_ := json.Marshal(req.Header)
	hader.UnmarshalJSON(hadbyte)
	proxyRequest.Headers = hader
	param :=  model.JSON{}
	param.UnmarshalJSON(p)
	proxyRequest.Params = param
	proxyRequest.Response = hader
	model.Create(proxyRequest)
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}