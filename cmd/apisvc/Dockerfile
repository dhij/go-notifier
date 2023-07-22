FROM golang:1.20-alpine AS builder

WORKDIR /build/

COPY . .
RUN go mod download

RUN SVC_ADDR=grpcsvc:9091 CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" github.com/dhij/go-notifier/cmd/apisvc

FROM scratch

WORKDIR /app/
ENV PATH=/app/bin/:$PATH

COPY --from=builder /build/apisvc ./bin/apisvc

EXPOSE 9090

CMD ["apisvc"]