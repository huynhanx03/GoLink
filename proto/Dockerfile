# Use a recent Go version that supports multi-arch
FROM golang:1.25-bookworm AS builder

# Install required tools
# We use 'go install' which compiles for the target architecture (including ARM64 for Mac) automatically.
ENV GO111MODULE=on

# Install buf
RUN go install github.com/bufbuild/buf/cmd/buf@v1.28.1

# Install protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.0

# Set up work directory
WORKDIR /workspace

# Copy script
COPY proto/gen.sh /usr/local/bin/gen-proto.sh
RUN chmod +x /usr/local/bin/gen-proto.sh

ENTRYPOINT ["/usr/local/bin/gen-proto.sh"]