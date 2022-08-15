FROM golang:alpine

RUN mkdir -p /app

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make

COPY . /app

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 3200

ENTRYPOINT ["./app"]