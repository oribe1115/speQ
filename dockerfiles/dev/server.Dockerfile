FROM golang:1.20.5-bullseye

WORKDIR /github.com/oribe1115/server/speQ

RUN apt install -y git \
    && go install github.com/cosmtrek/air@v1.44.0

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download