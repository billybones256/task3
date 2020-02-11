FROM golang

WORKDIR /go/src/grpc
COPY cmd/server cmd/server
COPY pkg pkg
WORKDIR /go/src/grpc/cmd/server

RUN go get -d -v
RUN go build

CMD  ./server
EXPOSE 8081
