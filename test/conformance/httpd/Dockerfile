# Base image and apt through nora (packages.maestra.io, 3-day hold, issues-maestra#1649);
# the PRIVATE_PACKAGES_TOKEN BuildKit secret is required.
FROM packages.maestra.io/docker-hub/library/debian:12-slim@sha256:734a72b3d6142a71993b4bc580ca01d8ca687f77a2c422d96a961d3376b179bc
RUN --mount=type=secret,id=PRIVATE_PACKAGES_TOKEN,required=true \
	--mount=type=bind,from=public.ecr.aws/g5f9s8a4/nora-apt:1.0.0@sha256:cd5fcec38f8b7c4fdd283d41f49ac656751d4deed168111f89f44976a43a8178,target=/nora-apt,rw \
	. /nora-apt/env.sh && apt-get update \
	&& apt-get install -y apache2 ssl-cert ca-certificates \
	&& apt-get clean
RUN \
	echo 'Listen 8443' > /etc/apache2/ports.conf \
	&& a2enmod headers proxy proxy_ajp proxy_http rewrite ssl \
	&& a2dissite 000-default.conf

COPY httpd/server.conf /etc/apache2/sites-enabled

COPY ssl/ory-conformity.crt /etc/ssl/certs/
COPY ssl/ory-conformity.key /etc/ssl/private/
COPY ssl/ory-conformity.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

ENTRYPOINT ["apachectl", "-DFOREGROUND"]
