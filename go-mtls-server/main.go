// code mainly ref https://smallstep.com/hello-mtls/doc/server/go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	CaCertPath     = "./tls/ca.pem"
	ServerCertPath = "./tls/server.pem"
	ServerKeyPath  = "./tls/server-key.pem"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(os.Stderr, "request from client: ua=%v ra=%v tls=%+v\n", req.UserAgent(), req.RemoteAddr, *req.TLS)
	fmt.Fprintf(w, "mtls works! now time: %s\n", time.Now().Local().Format(time.RFC3339))
}

func main() {
	caCert, _ := ioutil.ReadFile(CaCertPath)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// error example 1:
	// curl -k -vvv https://example.localhost:9443
	// tls: client didn't provide a certificate

	// error example 2:
	// curl  -vvv https://example.localhost:9443
	// local error: tls: bad record MAC

	// curl mtls the right way
	// 	curl --cert ./tls/client.pem --key ./tls/client-key.pem --cacert ./tls/ca.pem https://example.localhost:9443
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	http.HandleFunc("/", hello)

	server := &http.Server{
		Addr:      ":9443",
		TLSConfig: tlsConfig,
		Handler:   http.DefaultServeMux,
	}

	log.Println("https mtls server listen on: http://0.0.0.0:9443")
	log.Println(server.ListenAndServeTLS(ServerCertPath, ServerKeyPath))
}
