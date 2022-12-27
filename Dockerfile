FROM golang:1.19-alpine

RUN apk update && apk add git

WORKDIR /go/src/app 
COPY . .
RUN go get -v ./...
RUN go build

FROM alpine:latest
WORKDIR /home
COPY --from=0 /go/src/app/app .
CMD ["./app"]
