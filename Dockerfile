FROM golang:1.17-alpine3.15 as builder

COPY go.mod go.sum /build/
WORKDIR /build
RUN go mod download

COPY . /build
RUN go build -o service cmd/server/main.go

FROM alpine:3.15

RUN apk add tzdata

COPY --from=builder /build/service /app/
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "./service" ]
