FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata
ADD https://github.com/cloverzrg/file/raw/master/ca-certificates.crt /etc/ssl/certs/
ADD go-bin /app/go-bin
EXPOSE 9091
ENV HTTP_PORT 9091
ENV HTTP_LISTEN ''
ENV HTTP_EXTERNAL_URL ''
ENV CONSUL_ADDRESS 127.0.0.1:8085
ENV CONSUL_TOKEN ''
ENTRYPOINT ["/app/go-bin"]