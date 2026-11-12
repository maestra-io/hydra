# Packages, Go modules and base images only through nora (packages.maestra.io, 3-day hold,
# issues-maestra#1649), no public fallback: the PRIVATE_PACKAGES_TOKEN BuildKit secret is
# required (CI: the GitHub OIDC token; locally: your nra_ token). GOSUMDB stays default.
FROM packages.maestra.io/docker-hub/library/golang:1.26@sha256:978e0a4e3d953662577d784471db27742f543ab09eead06839a838188c73314b AS builder

RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,required=true \
    --mount=type=bind,from=public.ecr.aws/g5f9s8a4/nora-apt:1.0.0@sha256:cd5fcec38f8b7c4fdd283d41f49ac656751d4deed168111f89f44976a43a8178,target=/nora-apt,rw \
    . /nora-apt/env.sh && apt-get update && \
    apt-get install --no-install-recommends -y \
    git gcc bash ssl-cert ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/ory/hydra
RUN mkdir -p ./internal/httpclient

COPY oryx/go.mod oryx/go.mod
COPY oryx/go.sum oryx/go.sum
                

COPY go.mod go.sum ./
COPY internal/httpclient/go.* ./internal/httpclient/

ENV GO111MODULE on
ENV CGO_ENABLED 1

RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,env=NORA_TOKEN,required=true \
    GOPROXY="https://token:${NORA_TOKEN}@packages.maestra.io/go" go mod download

COPY . .

RUN go build  -o /usr/bin/hydra

VOLUME /mnt/sqlite

# Exposing the ory home directory
VOLUME /home/ory

# Declare the standard ports used by hydra (4444 for public service endpoint, 4445 for admin service endpoint)
EXPOSE 4444 4445

RUN mv test/conformance/ssl/ory-ca.* /etc/ssl/certs/ && \
    mv test/conformance/ssl/ory-conformity.crt /etc/ssl/certs/ && \
    mv test/conformance/ssl/ory-conformity.key /etc/ssl/private/ && \
    update-ca-certificates

ENTRYPOINT ["hydra"]
CMD ["serve"]
