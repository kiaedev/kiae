FROM golang:1.18-alpine as builder

RUN apk update && apk add curl just

WORKDIR /kiaedev/kiae
COPY . .
RUN just build


FROM debian:10

LABEL org.opencontainers.image.source=https://github.com/kiaedev/kiae

RUN apt-get update \
    && apt-get install -y ca-certificates telnet procps curl

ENV APP_HOME /srv
WORKDIR $APP_HOME

COPY --from=builder build/bin $APP_HOME/bin

CMD ["./bin/kiae", "server"]