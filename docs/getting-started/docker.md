# Docker

Run the readability CLI using Docker. The image is available from GitHub Container Registry.

!!! tip "Multi-Architecture Support"
    The image works on both Intel/AMD (`linux/amd64`) and ARM (`linux/arm64`) systems. Docker picks the right one for you.

## Quick Start

Pull the image and analyze your docs:

```bash
docker pull ghcr.io/adaptive-enforcement-lab/readability:latest

docker run --rm -v "$(pwd):/workspace" \
  ghcr.io/adaptive-enforcement-lab/readability:latest /workspace/docs
```

## Image Tags

Pick a tag based on how you want to track updates:

| Tag | Description |
|-----|-------------|
| `latest` | Most recent stable release |
| `vX.Y.Z` | Exact version (e.g., `v1.10.0`) |
| `vX.Y` | Latest patch for a minor version |
| `vX` | Latest release for a major version |

## Usage Examples

### Check with Thresholds

Fail if grade level is too high:

```bash
docker run --rm -v "$(pwd):/workspace" \
  ghcr.io/adaptive-enforcement-lab/readability:latest \
  --check --max-grade 12 /workspace/docs
```

### JSON Output

Get results as JSON for scripts or CI:

```bash
docker run --rm -v "$(pwd):/workspace" \
  ghcr.io/adaptive-enforcement-lab/readability:latest \
  --format json /workspace/docs
```

### Use a Config File

Mount your config and reference it:

```bash
docker run --rm -v "$(pwd):/workspace" \
  ghcr.io/adaptive-enforcement-lab/readability:latest \
  --config /workspace/.readability.yml /workspace/docs
```

## Security

!!! note "Signed Images"
    All images are signed with Cosign. You can verify them before use.

### Verify Image Signature

Check that the image came from our CI:

```bash
cosign verify ghcr.io/adaptive-enforcement-lab/readability:latest \
  --certificate-identity-regexp 'https://github.com/adaptive-enforcement-lab/readability/.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```

### Verify SBOM

Each image has a signed bill of materials (SBOM):

```bash
cosign verify-attestation ghcr.io/adaptive-enforcement-lab/readability:latest \
  --type cyclonedx \
  --certificate-identity-regexp 'https://github.com/adaptive-enforcement-lab/readability/.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```

### Base Image

The container uses Google's distroless base image. This means:

- No shell or package manager
- Runs as non-root user
- Only the readability binary is included

## CI/CD Examples

### GitLab CI

```yaml
readability:
  image: ghcr.io/adaptive-enforcement-lab/readability:latest
  script:
    - readability --check --max-grade 12 docs/
```

### CircleCI

```yaml
jobs:
  readability:
    docker:
      - image: ghcr.io/adaptive-enforcement-lab/readability:latest
    steps:
      - checkout
      - run: readability --check --max-grade 12 docs/
```

### Jenkins

```groovy
pipeline {
    agent {
        docker {
            image 'ghcr.io/adaptive-enforcement-lab/readability:latest'
        }
    }
    stages {
        stage('Check Readability') {
            steps {
                sh 'readability --check --max-grade 12 docs/'
            }
        }
    }
}
```
