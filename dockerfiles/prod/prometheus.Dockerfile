FROM alpine:3.18.2 as templete-build

WORKDIR /template

ARG SERVER_HOST

RUN echo $SERVER_HOST
RUN export SERVER_HOST

RUN apk add gettext
COPY config/prometheus/prometheus.yml.template .
RUN envsubst < prometheus.yml.template > prometheus.yml


FROM prom/prometheus:v2.45.0

COPY --from=templete-build /template/prometheus.yml  /etc/prometheus/prometheus.yml

EXPOSE 9090