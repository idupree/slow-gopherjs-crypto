package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func main() {
	// a 2048-bit RSA key
	keyfile := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA6iOnAfAuVuHp/DtWNvawWnx1FVeyJ2wg3NkzQMoU7SKkiLjB
u5fHVAL/5eBAt9K8xGMDTuglh7oxnRqNvosxmRZYahTb+3PxFzdmw4J0kvkbjnXO
lv/HDvABB/SKL8PUmUyuuf8qmM6b4yH6/jlfxOTTVricY1ztEMg+RKn6lzlnw2aB
1vbasiJdls2AX2YFM1KpeewXCla/NVTsw8z+DlC/QgCu4efpwNCxwshjmPox83cb
uKnlX7qfcldjxB+7zRnAwK8IxfEPlf6kYItH6s7vPRJwNKaXLpJOft5/z0lDPShb
bMhfG3djqmqIVUTbsiH5/3g5YwDqaHGaZJ6YpQIDAQABAoIBADuH2xOkFLQkdpCa
KZ6J9GbdVAucI4Dgfi6Izv4/Izjz4O+eQivHOHgE1zmEXhAe9b3dCt9KKRisX6b5
NGbpDzBEPTCjoxb48KidfnRhO/COg9Ih4+gCmoMSxMBYDbZ9TRW5KIfD9/moNnx2
jFGNFuuB3rYi9ChTXTcyIprkuQEvq833uuHlzDe65dSY/nY7XBGiz9gcN02mLyx0
uaJihN0BdFpl16keI6WdplYMrMu5ZTmSU2BGz8kp1izNqwzJPiqLJdQdq3lrJ8+N
Tb4tBICYaN3SSLuKp1IPxCT06GIfWS2zZqlsrcB8gJ5CdmubVgeDsIWcslFcFUkj
rrD+cwECgYEA+La9w813pa70z5NPOImUF/OmLm1X7QAVKB7DvAKngqxeJ4HKH2C7
8X8jWo7lGA16g+AKhrlsjMYtD3L/uvmmKiDFoMCOJFxOwbjAXSWUBn9hruBQTBe1
MI1us32bUT2QtG4kczQ1mj8vUYcpNqe8CmZGvVkytoRGRGpd97JZ4GUCgYEA8P+b
gFvdad6kXdOT3YB2vY+UnSwiXH7XZrbitQgPuT4tQpArOLO8V9IOvsWYECDriB0M
bGXRatoVwO9xmFB8Mmh6ppB1tFLq0NdvJQYqY8sgaXyO6TpXsUD2hrLknXYtsvFM
I+EyUKJL1VFYE788nuwN/zjZpWAVJOfsVq7us0ECgYEAubpdQ1awPn3EOy5aPnIe
sTQ1qP3mZwlkwy0WJdQlmyN0vDPj6EKoltLGZ93FnoySLOCle64ELavgAmVyKwI6
38LRUhX5D5YPCgMZQ3XIcIG3RIwl6mzZ1YQIYuktjyFWaibF+XHHhYQhpdQ1Jqdj
eLCfqs5tXHaysmxr94isLTECgYBY7Yj84WkfYtd2A9ehOYEXS78EDEAVr3xeW4fm
UzYyHb5nvHkmcDREl693N8R31x3yWP23lg6jhhvW2MQq9zANDb6MbevxAVPCgmxx
1geYUWEBa8P0TeID9zvA0oxHik8so5t79eIyHOEsstp0VvHQlrxHfDvbqUvEFYyB
JAXSwQKBgQDFvG669X09RXfC6XJTSSYuznTbhfuwGjI9hg6oEgYiluk+Q5cNQk8R
0iZhGfjNLcEFghjngwv8C0jjeEq2RaEsud2Cl8UMkm+f+mwvzMdWecYz75kGUXXO
6P3/YbCp9Ua2gN8dGTyga7lWE4p9WAxB/pnvj9Z/nOm/cWHAeHGl0A==
-----END RSA PRIVATE KEY-----
`)
	block, _ := pem.Decode(keyfile)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("parse error: ", err)
	}
	// sign an empty message
	hash := sha256.Sum256([]byte{})
	rsa.SignPKCS1v15(nil, key, crypto.SHA256, hash[:])
}
