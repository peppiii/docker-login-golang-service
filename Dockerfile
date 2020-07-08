FROM alpine:latest
RUN apk --no-cache --update add ca-certificates
RUN apk add --no-cache curl
RUN apk add --no-cache vim
COPY docker-login-service ./
ENTRYPOINT ["./docker-login-service"]
EXPOSE 8082
