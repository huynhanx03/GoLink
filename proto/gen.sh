#!/bin/bash
set -e

echo "Generating code from proto..."
# Ensure output directory exists
mkdir -p /workspace/Common/gen/go
mkdir -p /workspace/Common/gen/openapiv2

# Run buf generate
# We assume the container workdir is /workspace/proto
cd /workspace
buf generate proto --template proto/buf.gen.yaml
