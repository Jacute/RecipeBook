#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

DATABASE_HOSTNAME=127.0.0.1
MYSQL_ROOT_PASSWORD=redacted
MYSQL_DATABASE=redacted
MYSQL_PASSWORD=redacted
MYSQL_USER=redacted

cd src

echo -e "${GREEN}Generate test JWT keys...${NC}"
mkdir jwt_keys
../../gen_jwt_keys.sh

echo -e "${GREEN}Up database...${NC}"
docker compose run -d -p 3306:3306 recipebook_db
sleep 15

echo -e "${GREEN}Running...${NC}"
DATABASE_HOSTNAME=$DATABASE_HOSTNAME \
MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
MYSQL_DATABASE=$MYSQL_DATABASE \
MYSQL_PASSWORD=$MYSQL_PASSWORD \
MYSQL_USER=$MYSQL_USER \
JWT_PUBLIC_KEY_PATH=jwt_keys/public.pem go run .

echo -e "${RED}Down database...${NC}"
docker stop $(docker ps -lq)
rm -rf jwt_keys
