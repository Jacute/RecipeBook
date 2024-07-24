#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

export DATABASE_HOSTNAME=127.0.0.1
export MYSQL_ROOT_PASSWORD=test
export MYSQL_DATABASE=test
export MYSQL_PASSWORD=test
export MYSQL_USER=test

cd src

echo -e "${GREEN}Generate test JWT keys...${NC}"
mkdir jwt_keys
../../gen_jwt_keys.sh

echo -e "${GREEN}Up database...${NC}"
docker pull mysql:latest
docker run -d \
-e DATABASE_HOSTNAME=$DATABASE_HOSTNAME \
-e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
-e MYSQL_DATABASE=$MYSQL_DATABASE \
-e MYSQL_PASSWORD=$MYSQL_PASSWORD \
-e MYSQL_USER=$MYSQL_USER \
-p 127.0.0.1:3306:3306 mysql
sleep 15

export JWT_PUBLIC_KEY_PATH=../jwt_keys/public.pem
export JWT_PRIVATE_KEY_PATH=../jwt_keys/private.pem

echo -e "${GREEN}Running tests...${NC}"
go test tests/*.go

if [[ $? -eq 0 ]]; then
    echo -e "${GREEN}All tests passed${NC}"
else
    echo -e "${RED}Some tests failed${NC}"
fi

echo -e "${RED}Down database...${NC}"
docker stop $(docker ps -lq)
rm -rf jwt_keys
