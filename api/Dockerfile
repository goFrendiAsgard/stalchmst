FROM golang:1.13

WORKDIR /go/src/app
COPY . .

# we need netcat so that we can confirm mysql liveness using `wait.sh`
RUN apt-get update
RUN apt-get install netcat -y

# prepare go binary
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app

EXPOSE 3000

CMD ./wait.sh && ./app