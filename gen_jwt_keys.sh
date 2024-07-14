#!/bin/bash


openssl genrsa -out jwt_keys/private.key 2048
openssl rsa -in jwt_keys/private.key -pubout -out jwt_keys/public.key
