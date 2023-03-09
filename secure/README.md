# Create self-signed certificates with SANs
source:
- [Create self-signed certificates with Subject Alternative Names][https://www.youtube.com/watch?v=qoS4bLmstlk]
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
openssl req -new -nodes -out server.csr -keyout server.key -config localhost.csr.cnf
```

5. Create a server extension file (server_v3.ext). Example: https://www.mobilefish.com/download/openssl/sand.mobilefish_v3.ext.txt
- download the example
- modify the server extension file according to your situation
- add subject alternative name, e.g.
```

```
- In the server configuration file example (sand.mobilefish.cnf.example) I have used "CN=sand.mobilefish.com". 
This common name must be mentioned as one of the Subject Alternative Names (SANs).

6. The last step is to crete server certificate `server.crt` and serial number file `ca.srl`
```
sudo openssl x509 -req -in server.csr -CA ca.pem -CAkey privkey.pem  -CAcreateserial -out server.crt -days 3650  -extfile localhost_v3.ext 
```
Each issued certificate must contain a unique serial number assigned by the CA. It must be unique for each certificate given by a given CA. 
OpenSSL keeps the used serial numbers on a file.

# Server Certificate and Private Key
- The server certificate (server.crt) and server private key (server.key) are the two files you need to install on your server (Apache web server; proxy server)