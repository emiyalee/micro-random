FROM 172.16.3.221/public/golang:1.11 as builder
WORKDIR $GOPATH/src/github.com/emiyalee/random/random-service
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/random-service

FROM 172.16.3.221/public/centos:7.6.1810
WORKDIR /bin
COPY --from=builder /bin/random-service ./
CMD ["random-service", "-h"] 