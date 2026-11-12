# Packages, Go modules and base images only through nora (packages.maestra.io, 3-day hold,
# issues-maestra#1649), no public fallback: the PRIVATE_PACKAGES_TOKEN BuildKit secret is
# required (CI: the GitHub OIDC token; locally: your nra_ token). GOSUMDB stays default.
FROM packages.maestra.io/docker-hub/library/golang:1.26-alpine3.23@sha256:308364e4127bfa2b8d40676ef0b7627584f0f6f9e340f73929b01316fe35a222 AS builder

RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,required=true --mount=type=bind,from=public.ecr.aws/g5f9s8a4/nora-apk:1.0.0@sha256:55df38687c301e6fefa230dc1c71a8841618214cc549ade71be740b25822a00e,target=/nora-apk \
    /nora-apk/nora-apk -U --no-cache add build-base

WORKDIR /go/src/github.com/ory/x

ADD go.mod go.mod
ADD go.sum go.sum

RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,env=NORA_TOKEN,required=true \
    GOPROXY="https://token:${NORA_TOKEN}@packages.maestra.io/go" go mod download

ADD . .

RUN go build -o /usr/bin/eventlogger ./watcherx/integrationtest

FROM packages.maestra.io/docker-hub/library/alpine:3.22@sha256:49340b346798875b5592e5d420241f2f9cfba5f47e62ebb1446a1b34734e898f

COPY --from=builder /usr/bin/eventlogger /usr/bin/eventlogger

ENTRYPOINT ["eventlogger"]
CMD ["/etc/config/mock-config"]
