client:
	go run ./go-mtls-client


server:
	go run ./go-mtls-server


curl:
	curl --cert ./tls/client.pem --key ./tls/client-key.pem --cacert ./tls/ca.pem https://example.localhost:9443
