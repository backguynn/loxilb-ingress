FROM ghcr.io/loxilb-io/loxilb:latest

LABEL name="loxilb-ingress-manager" \
      vendor="loxilb.io" \
      version=$GIT_VERSION \
      release="0.1" \
      summary="loxilb-ingress-manager docker image" \
      description="ingress implementation for loxilb" \
      maintainer="backguyn@netlox.io"

WORKDIR /bin/
COPY ./bin/loxilb-ingress /bin/loxilb-ingress
USER root
RUN chmod +x /bin/loxilb-ingress
