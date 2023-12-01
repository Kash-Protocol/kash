# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/Kash-Protocol/kashd

WORKDIR /go/src/github.com/Kash-Protocol/kashd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache kashd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o kashd .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/Kash-Protocol/kashd/kashd /app/
COPY --from=build /go/src/github.com/Kash-Protocol/kashd/infrastructure/config/sample-kashd.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
