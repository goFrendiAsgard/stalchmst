# tear down
docker-compose down

# setup
docker-compose up -d

# wait
env MYSQL_HOST=localhost MYSQL_PORT=3307 ./wait.sh

# test
env MYSQL_CONNECTION_STRING="root:toor@tcp(localhost:3306)/template" go test