package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"io/ioutil"
	"log"
)

type Pxy struct {}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s %s\n", req.Method, req.Host, req.RemoteAddr,req.URL.String())

	transport :=  http.DefaultTransport

	// step 1
	outReq := new(http.Request)
	*outReq = *req // this only does shallow copies of maps

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	// step 2
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}

	rw.WriteHeader(res.StatusCode)
	by,_ := ioutil.ReadAll(res.Body)
	log.Print(string(by))
	io.Copy(rw, res.Body)
	res.Body.Close()
}

func main() {
	fmt.Println("Serve on :8081")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8081", nil)
}