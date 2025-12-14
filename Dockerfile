# Uses pre-built binaries from CI - no build stage needed
# Binary is selected based on TARGETARCH (amd64 or arm64)
FROM gcr.io/distroless/static-debian12:nonroot@sha256:2b7c93f6d6648c11f0e80a48558c8f77885eb0445213b8e69a6a0d7c89fc6ae4

ARG TARGETARCH
COPY dist/readability_linux_${TARGETARCH} /usr/local/bin/readability

ENTRYPOINT ["readability"]
