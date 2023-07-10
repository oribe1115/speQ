FROM node:18.16-alpine3.18 as client-build

WORKDIR /app

COPY client/package*.json .
RUN npm ci

COPY client/ .

RUN npm run build