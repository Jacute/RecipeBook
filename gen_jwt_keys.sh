#!/bin/bash


openssl genrsa -out jwt_keys/private.pem 2048
openssl rsa -in jwt_keys/private.pem -pubout -out jwt_keys/public.pem
