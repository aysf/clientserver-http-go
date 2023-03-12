# Some examples of client-server application

1. Client-Server simple

Learn how to make client-server very simple in Go

2. Client-Server With Mux. learn how to use mux.

open terminal and run `make server-mux`, and open another terminal and run `make client-mux`

3. Client-Server With Mux and Var

Learn how to get variable from url path and query. To use it, open two terminals in the same time to server `make server-mux-var` and client `make client-mux-var`. 

4. Client-Server With Mux and Post

Learn how to work with Post request, header request, Json, and form. To use it, open two terminals in the same time to server `make server-mux-post` and client `make client-mux-post`. 

5. Client-Server insecure

Learn how to use cert and private key in server. Prior to start, generate server private key and certificate (see readme). To start, open two terminals in the same time to server `make server-secure` and client `make client-insecure`. But in this scenario, the request from client is still insecure. 

To get secure connection, there are two ways:
- Add cert into certPool in client code (see the next step 6)
- Add cert into local system (see this article https://github.com/aysf/clientserver-http-go/tree/main/secure#how-to-add-certificate-in-system)

6. Client-Server secure

Learn how to make secure connection for both client and server with self-signed cert and SANs. Now we should generate again private key and certificate key but with SANs configuration and then start by opening two terminal for running server `make server-secure` and client `make client-secure`

7. Client-Server http2

Learn how to use http2 for client-server connection

8. Client-Server with PSK

Learn how to build client-server with pre-shared key