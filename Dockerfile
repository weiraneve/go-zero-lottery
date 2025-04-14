FROM golang:1.24 AS builder

LABEL stage=gobuilder
ENV CGO_ENABLED=1
WORKDIR /build

ENV GOPROXY=https://goproxy.cn,direct

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY etc /app/etc
RUN go build -ldflags="-s -w" -o /app/lottery ./lottery.go

FROM frolvlad/alpine-glibc:latest

WORKDIR /app
COPY --from=builder /app/lottery /app/lottery
COPY --from=builder /app/etc /app/etc

CMD ["./lottery", "-f", "etc/lottery.yaml"]
