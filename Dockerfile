FROM golang:1.21 AS build-env

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test ./...

RUN mkdir -p /bin \
	&& GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o ./bin/app ./cmd

FROM alpine:3
WORKDIR /app
COPY --from=build-env /app .
CMD ["./bin/app"]
