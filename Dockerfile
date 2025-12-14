# Uses pre-built binaries from CI - no build stage needed
# Binary is selected based on TARGETARCH (amd64 or arm64)
FROM gcr.io/distroless/static-debian12:nonroot

ARG TARGETARCH
COPY dist/readability_linux_${TARGETARCH} /usr/local/bin/readability

ENTRYPOINT ["readability"]
