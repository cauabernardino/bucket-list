FROM golang:1.16.4-alpine3.13 as builder
RUN mkdir /build
ADD go.mod go.sum /build/
WORKDIR /build
RUN go mod download
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/bucket-list github.com/cauabernardino/bucket-list
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine
WORKDIR /app
RUN apk add --no-cache ca-certificates && update ca-certificates
COPY --from=builder /build/bin/bucket-list .
COPY --from=builder /build/migrate.linux-amd64 ./migrate
EXPOSE 8080
ENTRYPOINT ["/app/bucket-list"]