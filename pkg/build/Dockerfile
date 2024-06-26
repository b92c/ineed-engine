# Etapa de construção
FROM public.ecr.aws/docker/library/golang:1.23rc1-alpine3.20 AS build

WORKDIR /go/src

RUN apk update && apk --no-cache add gcc musl-dev make git

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN go install github.com/cosmtrek/air@v1.49.0
RUN mkdir -p bin && go build -o bin/main ./cmd/api

CMD ["air", "-c", ".air.toml"]
