ARG KASHD_IMAGE
ARG KASHMINER_IMAGE

FROM ${KASHD_IMAGE} as kashd
FROM ${KASHMINER_IMAGE} as kashminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/Kash-Protocol/kashd

WORKDIR /go/src/github.com/Kash-Protocol/kashd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=kashd /app/ /app/
COPY --from=kashminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/Kash-Protocol/kashd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]