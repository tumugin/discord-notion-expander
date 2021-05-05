FROM golang:1.16.3 AS builder
WORKDIR /go/src/app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
RUN go build

FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates
RUN useradd -m app
USER app
WORKDIR /app
COPY --from=builder --chown=app:app /go/src/app/discord-notion-expander .
CMD ["./discord-notion-expander"]
