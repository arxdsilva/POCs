FROM golang:1.16.5-buster as builder
LABEL stage=intermediate

ARG GITHUB_PRIVATE_KEY
ARG GOPROXY
ARG DOCKER_COMPOSE
ENV GO111MODULE=on

WORKDIR /app
COPY . ./
RUN go build -a -o ./bin/svc

FROM debian:buster-slim
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/svc /svc

RUN mkdir /data
ADD ./default.pem /data/default.pem
VOLUME /data

WORKDIR /

CMD ["./svc"]
