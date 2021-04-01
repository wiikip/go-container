FROM golang:alpine AS building

WORKDIR /app

COPY . /app

 
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/go-server
FROM scratch

COPY --from=building /go/bin/go-server /go/bin/go-server


ENTRYPOINT ["/go/bin/go-server"]

