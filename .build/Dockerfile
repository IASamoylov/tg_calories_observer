ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS build

ARG APP_LDFLAGS

WORKDIR /app

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN CGO_ENABLED=0 go build -ldflags="${APP_LDFLAGS}" -o=./build/calories-observer-telegram-bot ./cmd/telegram/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=build  /app/build/calories-observer-telegram-bot .

CMD ["./calories-observer-telegram-bot"]