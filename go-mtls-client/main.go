package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	CaCertPath     = "./tls/ca.pem"
	ClientCertPath = "./tls/client.pem"
	ClientKeyPath  = "./tls/client-key.pem"
)

func main() {
	caCert, err := os.ReadFile(CaCertPath)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(ClientCertPath, ClientKeyPath)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	for {
		mtlsClientRequest(client)
		time.Sleep(time.Second * 2)
	}
}

func mtlsClientRequest(client *http.Client) {
	// Make a request
	mtlsReqUrl := "https://example.localhost:9443"
	r, err := client.Get(mtlsReqUrl)
	if err != nil {
		log.Printf("request failed, url=%v err=%v", mtlsReqUrl, err)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body failed, url=%v err=%v", mtlsReqUrl, err)
		return
	}
	log.Printf("request successfully, url=%v body=%s", mtlsReqUrl, body)
}
