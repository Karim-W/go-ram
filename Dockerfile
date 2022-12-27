FROM golang:1.19-alpine

RUN apk update && apk add git

WORKDIR /
COPY . .
RUN go mod tidy
RUN go build -o /app /main.go

FROM alpine:latest
WORKDIR /home
COPY --from=0 /app /app

EXPOSE 8080
CMD ["./app"]
