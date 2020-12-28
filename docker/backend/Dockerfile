# ------
FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

RUN apk update \
  && apk add make git

ADD . .

RUN make fix

ENV GOARCH=arm
RUN make build

# ------
FROM golang:1.14-alpine as final

COPY --from=builder /app/build/freegrow /usr/local/bin/freegrow

WORKDIR /freegrow

ENTRYPOINT [ "/usr/local/bin/freegrow" ]
