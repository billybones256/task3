FROM golang

WORKDIR /go/src/task3
COPY cmd/server cmd/server
COPY pkg pkg
WORKDIR /go/src/task3/cmd/server

RUN go get -d -v
RUN go build

CMD  ./server
EXPOSE 8081
