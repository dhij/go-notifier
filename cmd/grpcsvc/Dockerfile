FROM golang:1.20-alpine AS builder
ENV DB_SOURCE "root:password@tcp(mysql:3306)/notifier_db"

WORKDIR /build/

COPY . .
RUN go mod download

RUN DB_SOURCE=${DB_SOURCE} CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" github.com/dhij/go-notifier/cmd/grpcsvc

RUN CGO_ENABLED=0 GOOS=linux go install -a -ldflags "-s -w" -tags "mysql" \
    github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

FROM scratch

WORKDIR /app/
ENV PATH=/app/bin/:$PATH

COPY --from=builder /build/grpcsvc ./bin/grpcsvc
COPY --from=builder /go/bin/migrate ./bin/migrate
COPY --from=builder /build/db/ .

EXPOSE 9091

CMD ["grpcsvc"]