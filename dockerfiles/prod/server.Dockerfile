FROM golang:1.20.5-alpine3.18

WORKDIR /github.com/oribe1115/server/speQ

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

COPY server/ .

RUN go build -o app .

ENTRYPOINT [ "./app" ]