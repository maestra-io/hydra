# Base image, Maven and apt through nora (packages.maestra.io, 3-day hold, issues-maestra#1649);
# the PRIVATE_PACKAGES_TOKEN BuildKit secret is required.
FROM packages.maestra.io/docker-hub/library/maven:3-jdk-11@sha256:805f366910aea2a91ed263654d23df58bd239f218b2f9562ff51305be81fa215

WORKDIR /usr/src/mymaven

RUN wget https://gitlab.com/openid/conformance-suite/-/archive/release-v4.1.4/conformance-suite-release-v4.1.4.zip && \
  unzip conformance-suite-release-v4.1.4.zip -d . && \
  rm conformance-suite-release-v4.1.4.zip && \
  find conformance-suite-release-v4.1.4 -maxdepth 1 -mindepth 1 -exec mv {} . \; && \
  rmdir conformance-suite-release-v4.1.4

COPY <<'EOF' /root/.m2/settings.xml
<settings>
  <mirrors><mirror><id>nora</id><mirrorOf>*</mirrorOf><url>https://packages.maestra.io/maven2</url></mirror></mirrors>
  <servers><server><id>nora</id><username>token</username><password>${env.NORA_TOKEN}</password></server></servers>
</settings>
EOF

RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,required=true \
    --mount=type=bind,from=public.ecr.aws/g5f9s8a4/nora-apt:1.0.0@sha256:cd5fcec38f8b7c4fdd283d41f49ac656751d4deed168111f89f44976a43a8178,target=/nora-apt,rw \
    NORA_TOKEN="$(cat /run/secrets/PRIVATE_PACKAGES_TOKEN)" mvn -B clean package -DskipTests && \
    . /nora-apt/env.sh && apt-get update && apt-get install -y \
    redir ca-certificates

COPY ssl/ory-conformity.crt /etc/ssl/certs/
COPY ssl/ory-conformity.key /etc/ssl/private/
COPY ssl/ory-conformity.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

CMD java -Xdebug -Xrunjdwp:transport=dt_socket,address=*:9999,server=y,suspend=n -jar /usr/src/mymaven/target/fapi-test-suite.jar --fintechlabs.base_url=https://httpd:8443 --fintechlabs.devmode=true --fintechlabs.startredir=true
