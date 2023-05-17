ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS build

ARG APP_LDFLAGS

WORKDIR /app

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="${APP_LDFLAGS}" -o=./build/tbco-web-server ./cmd/telegram/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=build  /app/build/tbco-web-server /

EXPOSE 9090

CMD ["/tbco-web-server"]