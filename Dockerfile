FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src/backend

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .
RUN go build -o ./bin cmd/app/main.go

FROM alpine

WORKDIR /usr/local/src/backend

COPY --from=builder /usr/local/src/backend/bin /usr/local/src/backend/app
COPY configs/config.yml configs/config.yml

CMD ["/usr/local/src/backend/app"]