Generate TLS certificate for server

1. Install OpenSSL: If you don't already have OpenSSL installed on your system, you'll need to install it first. OpenSSL is available for most operating systems and can usually be installed using your system's package manager.
```sh
openssl genrsa -out server.key 2048
```

2. Generate a private key: Use the `openssl` command to generate a private key:
```sh
openssl req -new -key server.key -out server.csr
```
This will generate a 2048-bit RSA private key and save it to a file named `server.key`.

3. Create a certificate signing request (CSR): Use the `openssl` command to create a certificate signing request:
```sh
openssl req -new -key server.key -out server.csr
```
This will create a new CSR using the private key you generated in step 2. You'll be prompted to enter some information about the certificate, such as the Common Name (CN) for the server (e.g. "localhost").

4. Generate a self-signed certificate: Use the `openssl` command to generate a self-signed certificate:
```sh
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```
This will generate a self-signed TLS certificate that is valid for 365 days and save it to a file named `server.crt`.

Note that since this is a self-signed certificate, it will not be trusted by default by most web browsers or operating systems. You'll likely need to add an exception or import the certificate into your trust store in order to use it.

For generating client certificates and keys for testing purposes, follow these steps:

1. Generate a client key:
```sh
openssl genpkey -algorithm RSA -out client.key
```

2. Generate a certificate signing request (CSR) for the client key:
```sh
openssl req -new -key client.key -out client.csr
```

3. Generate a client certificate using the server's CA key and certificate:
```sh
openssl x509 -req -in client.csr -CA server.crt -CAkey server.key -CAcreateserial -out client.crt -days 365
```

Note that in step 3, we are using the server's CA key and certificate to sign the client certificate. This assumes that the server is acting as a certificate authority (CA) and has already generated a CA key and certificate. If you don't have a CA key and certificate, you can generate them using the same openssl commands that were used to generate the server key and certificate. However, it's important to note that in a production environment, you would typically obtain client certificates from a trusted CA, rather than generating them yourself.
