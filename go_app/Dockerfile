FROM golang:1.20-alpine3.18 AS builder
RUN apk --no-cache add bash && \
    mkdir src/go_app && \
    mkdir src/go_app/cmd && \
    mkdir src/go_app/pkg && \
    mkdir src/go_app/scripts && \
    mkdir src/go_app/sql
WORKDIR /go/src/go_app 
COPY go.mod go.sum ./
COPY pkg ./pkg
COPY cmd ./cmd
COPY scripts ./scripts
COPY sql ./sql

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./bin/go_app ./cmd/taskmanager/start.go

FROM alpine:3.18
RUN apk --no-cache add bash curl && \
    mkdir /go && \
    mkdir /go/bin && \
    mkdir /go/bin/go_app

WORKDIR /go/bin/go_app
COPY --from=builder /go/src/go_app/bin/* /go/bin/go_app/
CMD ["./go_app"]
EXPOSE 8080