#!/bin/sh

# setup
echo ">> Setting up testing environment"
docker-compose down
docker-compose up -d

# waiting
echo ">> Wait MySQL to be ready"
while ! docker exec stalchmst-mysql-test mysql -u user -ppassword -e "USE db;"; do
  sleep 1 # wait for 1 second to try again
  echo ">> Retry..."
done

# test
echo ">> Start testing"
env MYSQL_CONNECTION_STRING="user:password@tcp(localhost:3307)/db" MYSQL_MAX_CONNECT_ATTEMPT=50 go test

# tear down
echo ">> Clean up testing environment"
docker-compose down
echo ">> Done"
