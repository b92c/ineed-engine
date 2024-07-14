FROM golang:1.23rc1-alpine3.20 

WORKDIR /go/src

RUN apk update && apk --no-cache add gcc musl-dev make git

COPY . .

RUN go install github.com/air-verse/air@latest

RUN make build


CMD ["air", "-c", ".air.toml"]