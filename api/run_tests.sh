#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

DATABASE_HOSTNAME=127.0.0.1
MYSQL_ROOT_PASSWORD=test
MYSQL_DATABASE=test
MYSQL_PASSWORD=test
MYSQL_USER=test


echo -e "${GREEN}Up database...${NC}"
docker pull mysql:latest
docker run -d \
-e DATABASE_HOSTNAME=$DATABASE_HOSTNAME \
-e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
-e MYSQL_DATABASE=$MYSQL_DATABASE \
-e MYSQL_PASSWORD=$MYSQL_PASSWORD \
-e MYSQL_USER=$MYSQL_USER \
-p 127.0.0.1:3306:3306 mysql
sleep 12

echo -e "${GREEN}Running test...${NC}"
DATABASE_HOSTNAME=$DATABASE_HOSTNAME \
MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
MYSQL_DATABASE=$MYSQL_DATABASE \
MYSQL_PASSWORD=$MYSQL_PASSWORD \
MYSQL_USER=$MYSQL_USER go test tests/*

if [[ $? -eq 0 ]]; then
    echo -e "${GREEN}All tests passed${NC}"
else
    echo -e "${RED}Some tests failed${NC}"
fi

echo -e "${RED}Down database...${NC}"
docker stop $(docker ps -lq)
