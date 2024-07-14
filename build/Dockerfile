FROM golang:1.23rc1-alpine3.20 AS build

WORKDIR /go/src

RUN apk update && apk --no-cache add gcc musl-dev make git

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN mkdir -p bin && go build -o bin/main ./cmd/api

FROM golang:1.23rc1-alpine3.20

WORKDIR /go/src

RUN apk update && apk --no-cache add gcc musl-dev make git

COPY --from=build /go/src/bin /go/src/bin
COPY --from=build /go/src/go.mod /go/src/go.mod
COPY --from=build /go/src/go.sum /go/src/go.sum
COPY . .

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
