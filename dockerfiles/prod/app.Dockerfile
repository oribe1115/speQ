# クライアントサイドのビルド
FROM node:18.16-alpine3.18 as client-build

WORKDIR /app

COPY client/package*.json .
RUN npm ci

COPY client/ .

RUN npm run build


# サーバーサイドのビルド
FROM golang:1.20.5-alpine3.18 as server-build

WORKDIR /github.com/oribe1115/server/speQ

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

COPY server/ .

RUN go build -o app .


# 最終的な配信用
FROM nginx:1.25.1-alpine

COPY /config/nginx/server.conf /etc/nginx/conf.d/default.conf

COPY --from=client-build /app/dist /usr/share/nginx/html
COPY --from=server-build /github.com/oribe1115/server/speQ/app /server/app

EXPOSE 80

ENTRYPOINT [ "/server/app" ]