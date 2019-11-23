FROM golang:1.13 AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./...

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=build /app/main /usr/bin/main

ENTRYPOINT [ "main" ]
CMD [ "--help" ]
