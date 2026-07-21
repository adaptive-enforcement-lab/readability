# Uses pre-built binaries from CI - no build stage needed
# Binary is selected based on TARGETARCH (amd64 or arm64)
FROM gcr.io/distroless/static-debian12:nonroot@sha256:f5b485ea962d9bd1186b2f6b3a061191539b905b82ec395de78cbfae51f20e35

ARG TARGETARCH
COPY dist/readability_linux_${TARGETARCH} /usr/local/bin/readability

ENTRYPOINT ["readability"]
