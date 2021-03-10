FROM golang:1.15-alpine
COPY . .
ENV GOPATH ""
RUN mkdir -p /go/bin
RUN go build -o /go/bin/qotd
ENTRYPOINT ["/go/bin/qotd"]
