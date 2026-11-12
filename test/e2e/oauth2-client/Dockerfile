# Base image and npm through nora (packages.maestra.io, 3-day hold, issues-maestra#1649); the
# PRIVATE_PACKAGES_TOKEN BuildKit secret is required and goes into a tmpfs npmrc, never a layer.
FROM packages.maestra.io/docker-hub/library/node:10.8-alpine@sha256:962dc023d8598778df81a1a52719bd4ed7f341661742330bfdc72ca8241e58d6

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app
RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,required=true --mount=type=tmpfs,target=/npm-auth \
    printf 'registry=https://packages.maestra.io/npm/\n//packages.maestra.io/npm/:_authToken=%s\n' "$(cat /run/secrets/PRIVATE_PACKAGES_TOKEN)" > /npm-auth/npmrc && \
    NPM_CONFIG_USERCONFIG=/npm-auth/npmrc npm install --silent; exit 0

ENTRYPOINT npm start

EXPOSE 3000