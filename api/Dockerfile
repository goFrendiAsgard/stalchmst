FROM golang:1.13

WORKDIR /go/src/app
COPY . .

# prepare go binary
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app

EXPOSE 3000

CMD ./app