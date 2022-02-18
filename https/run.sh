#!/usr/bin/env bash
openssl req -new -nodes -x509 -out ./server.crt -keyout ./server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=yangjin94@163.com"
go run server-tls.go

