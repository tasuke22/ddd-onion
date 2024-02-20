FROM golang:1.22.0-alpine3.18

WORKDIR /data
COPY ./ ./

WORKDIR /data/app

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install go.uber.org/mock/mockgen@v0.3.0 && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]