#!/bin/sh

echo "Waiting MySQL ${MYSQL_HOST}:${MYSQL_PORT}"
while ! nc -z ${MYSQL_HOST} ${MYSQL_PORT}; do
  sleep 0.1 # wait for 1/10 of the second before check again
done

echo "MySQL ready"
