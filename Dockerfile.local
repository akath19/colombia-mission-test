#Builder container
FROM golang:buster as builder
WORKDIR /go/app
COPY . .
ENV CGO_ENABLED=0
RUN go get -d -v github.com/lib/pq
RUN go build -o program -v ./...

#Runner container
FROM alpine
WORKDIR /app
COPY --from=builder /go/app/ .
#Add wait script to wait for PostgreSQL before starting
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.6.0/wait /wait
RUN chmod +x /wait
CMD /wait && /app/program