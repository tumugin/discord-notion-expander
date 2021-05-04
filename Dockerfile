FROM golang:1.16.3 AS builder

WORKDIR /go/src/app
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .
RUN go build

FROM golang:1.16.3
RUN useradd -m app
USER app
WORKDIR /app
COPY --from=builder --chown=app:app /go/src/app/discord-notion-expander .
RUN chmod +x ./discord-notion-expander
CMD ["./discord-notion-expander"]
