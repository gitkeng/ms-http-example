#!/bin/bash
FILE_CERT_NAME=localhost
openssl req -x509 \
            -sha512 -days 18250 \
            -nodes \
            -newkey rsa:4096 \
            -subj "/CN=selfsigned.app/C=TH/L=Bangkok" \
            -addext "subjectAltName = DNS:selfsigned.app" \
            -keyout "$FILE_CERT_NAME.key" -out "$FILE_CERT_NAME.crt"
