#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD app app
ENTRYPOINT ["/app"]
