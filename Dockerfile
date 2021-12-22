FROM golang:alpine as build

RUN apk --no-cache add ca-certificates

FROM alpine:latest

RUN apk --no-cache add postgresql-client

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY script/wait-for-postgres.sh /wait-for-postgres.sh

RUN chmod +x wait-for-postgres.sh

COPY release/main /

CMD ["/main"]