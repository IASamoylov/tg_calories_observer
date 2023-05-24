ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS build

ARG APP_LDFLAGS

WORKDIR /app

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="${APP_LDFLAGS}" -o=./build/calories-observer-telegram-bot ./cmd/telegram/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=build  /app/build/calories-observer-telegram-bot /

EXPOSE 9090

CMD ["/calories-observer-telegram-bot"]