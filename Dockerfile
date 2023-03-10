FROM golang:1.19-alpine

RUN apk update && apk add git

WORKDIR /
COPY . .

RUN go build -o /app .

FROM alpine:latest
WORKDIR /
COPY --from=0 /app /app

EXPOSE 8080
CMD ["/app"]
