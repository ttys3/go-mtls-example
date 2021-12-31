# go-mtls-example

a simple https client and server example to demonstrate the Mutual TLS authentication

## What is mutual TLS (mTLS)?

ref: https://www.cloudflare.com/learning/access-management/what-is-mutual-tls/

**Mutual TLS**, or **mTLS** for short, is a method for mutual authentication. mTLS ensures that the parties at each end of a network connection are who they claim to be by verifying that they both have the correct private key. The information within their respective TLS certificates provides additional verification.

mTLS is often used in a Zero Trust security framework* to verify users, devices, and servers within an organization. It can also help keep APIs secure.

*Zero Trust means that no user, device, or network traffic is trusted by default, an approach that helps eliminate many security vulnerabilities.

## Refs

A Complete Guide to Securely Connecting Go Servers and Clients Using Mutual TLS https://smallstep.com/hello-mtls/doc/combined/go/go

Configuring Your Go Server for Mutual TLS https://smallstep.com/hello-mtls/doc/server/go

Generating self-signed root CA certificate and private key
https://github.com/cloudflare/cfssl#generating-self-signed-root-ca-certificate-and-private-key


Generating a local-issued certificate and private key
https://github.com/cloudflare/cfssl#generating-a-local-issued-certificate-and-private-key

mTLS on production usage: https://learn.hashicorp.com/tutorials/nomad/security-enable-tls#node-certificates
