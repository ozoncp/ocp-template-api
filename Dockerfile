# Builder

ARG GITHUB_PATH=github.com/ozoncp/ocp-template-api

FROM golang:1.16-alpine AS builder
RUN apk add --update make git protoc protobuf protobuf-dev curl
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
RUN make deps && make build

RUN go get github.com/go-delve/delve/cmd/dlv

# gRPC Server

FROM alpine:latest as server
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/grpc-server .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

COPY --from=builder /go/bin/dlv .

RUN chown root:root grpc-server

EXPOSE 50051
EXPOSE 40000
EXPOSE 8080
EXPOSE 9100


# CMD ["./grpc-server", "--migration", "up"]
CMD ["./dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./grpc-server"]
