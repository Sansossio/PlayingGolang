FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v "./src/"
RUN go install -v "./src/main.go"

CMD ["app"]