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
CMD /app/program