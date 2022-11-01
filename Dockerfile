FROM golang:1.18-bullseye as builder

RUN curl -q 'https://proget.makedeb.org/debian-feeds/prebuilt-mpr.pub' | gpg --dearmor | tee /usr/share/keyrings/prebuilt-mpr-archive-keyring.gpg 1> /dev/null \
    && echo "deb [signed-by=/usr/share/keyrings/prebuilt-mpr-archive-keyring.gpg] https://proget.makedeb.org prebuilt-mpr bullseye" | tee /etc/apt/sources.list.d/prebuilt-mpr.list \
    && apt-get update \
    && apt-get install -y just

WORKDIR /kiaedev/kiae
COPY . .
RUN just build


FROM debian:bullseye

LABEL org.opencontainers.image.source=https://github.com/kiaedev/kiae

RUN apt-get update \
    && apt-get install -y ca-certificates telnet procps curl

ENV APP_HOME /srv
WORKDIR $APP_HOME

COPY --from=builder /kiaedev/kiae/build/bin $APP_HOME/bin

CMD ["./bin/kiae", "server"]