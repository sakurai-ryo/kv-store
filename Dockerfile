FROM golang:1.17-alpine as builder

ENV ROOT=/go/src/kv-store

WORKDIR ${ROOT}

COPY . ${ROOT}

RUN go mod download

RUN GOOS=linux GOARCH=arm64 go build -o ${ROOT}/bin/kv-store ./main.go


FROM alpine:3.15 as prod

ENV ROOT=/go/src/kv-store

WORKDIR ${ROOT}

COPY --from=builder ${ROOT}/bin/kv-store ${ROOT}

EXPOSE 8080

CMD ["/go/src/kv-store/kv-store"]
