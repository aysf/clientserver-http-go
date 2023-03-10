# How to create certificate

**Server Certificate and Private Key**

The server certificate (server.crt) and server private key (server.key) are the two files you need to install on your server (Apache web server; proxy server)

## Create self-signed certificates (old way)
source: https://dasarpemrogramangolang.novalagung.com/C-https-tls.html

1. Generate private key, using this command:
```
openssl genrsa -out server.key 2048
```
and this command:
```
openssl ecparam -genkey -name secp384r1 -out server.key
```

2. Generate self-signed certificate (that contains public key) from the private key 
```
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

3. Fill the form in terminal. In **Common Name** put with `localhost` for local testing or any domains. 

## Create self-signed certificates (new)
source: 
- http://www.cs.toronto.edu/~arnold/427/19s/427_19S/tool/ssl/notes.pdf
- https://www.feistyduck.com/library/openssl-cookbook/

1. generate private key for server `server.key`
```
openssl genpkey -out server.key -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -pkeyopt rsa_keygen_pubexp:3
```
2. generate self-signed certificate for server
```
openssl req -new -x509 -days 365 -key server.key -out server.crt
```

## Create self-signed certificates with SANs
source:
- [Create self-signed certificates with Subject Alternative Names](https://www.youtube.com/watch?v=qoS4bLmstlk)
- [Key Players of SSL & TLS: Client, Server, Certificate Authority (CA) - Practical TLS](https://www.youtube.com/watch?v=C7Y4UEBJ0Og)
1. Create a 2048 bit CA private key:
```
openssl genrsa -out privkey.pem 2048
```
2. Create a self-signed CA certificate:
```
openssl req -new -x509 -days 365 -nodes -key privkey.pem -sha256 -out ca.pem
```
with the options as follows:
```
Country Name: NL
State or Province Name: Noord-Holland
Locality Name: Zaandam
Organization Name: Mobilefish.com CA
```
the suffix `CA` after `Mobilefish.com` means 'Certificate Authority'. CA is the governing entity that issues certificates to servers. It is trusted by both the client and the server, and also provides what's known as the trust anchor. The idea is the client might not trust all servers innately but if the client trust the CA and the CA provided an identity a certificate for the server therefore the client can trust the server. That is what the trust anchor role is for the certificate authority.    

As of today five organizations secure 98% of the internet. Here are those five organizations:
- IdenTrust (51,9%)
- DigiCert (19,4%)
- Sectigo (17,5%)
- GoDaddy (6,9%)
- GlobalSign (2,9%)

3. Create server configuration file
- create a server configuration file (localhost.csr.cnf). Example https://www.mobilefish.com/download/openssl/sand.mobilefish.csr.cnf.txt
- modify the server configuration file according to your situation

4. Create a server Certificate Signing Request (CSR) and server private key
```
openssl req -new -nodes -out server.csr -keyout server.key -config localhost.cnf
```

5. Create a server extension file (server_v3.ext). Example: https://www.mobilefish.com/download/openssl/sand.mobilefish_v3.ext.txt
- download the example
- modify the server extension file according to your situation
- add subject alternative name, e.g.
```
[alt_names]
DNS.1 = sand.mobilefish.com 
DNS.2 = proxy.mobilefish.com 
```
- In the server configuration file example (sand.mobilefish.cnf.example) I have used "CN=sand.mobilefish.com". 
This common name must be mentioned as one of the Subject Alternative Names (SANs).

6. The last step is to crete server certificate `server.crt` and serial number file `ca.srl`
```
openssl x509 -req -in server.csr -CA ca.pem -CAkey privkey.pem  -CAcreateserial -out server.crt -days 3650  -extfile localhost_v3.ext 
```
Each issued certificate must contain a unique serial number assigned by the CA. It must be unique for each certificate given by a given CA. 
OpenSSL keeps the used serial numbers on a file.

## Create self-signed certificates with SANs (new)

1. generate private key

```
openssl genpkey -out server.key -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -pkeyopt rsa_keygen_pubexp:3
```

2. create cofig file `*.cnf`

```
[req]
prompt = no
distinguished_name = dn
req_extensions = ext
input_password = PASSPHRASE
[dn]
CN = www.feistyduck.com
emailAddress = ansufw@gmail.com
O = Feisty Duck Ltd
L = London
C = GB
[ext]
subjectAltName = DNS:www.feistyduck.com,DNS:feistyduck.com
```
change the domain names above to localhost for testing in local. see `server.cnf` as example.

3. generate certificate sign request `*.csr` file
```
openssl req -new -config server.cnf -key server.key -out server.csr   
```

4. create extension file `*.ext`
create file ext, e.g. `server.ext` with list the desired hostnames
```
subjectAltName = DNS:*.feistyduck.com, DNS:feistyduck.com
```
change the domain name to `localhost` for testing in local.

5. generate the `*.crt` certificate file
```
openssl x509 -req -days 365 -in server.csr -signkey server.key -extfile server.ext -out server.crt 
```