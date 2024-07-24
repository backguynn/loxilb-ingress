FROM golang:1.22-alpine

LABEL name="loxilb-ingress-manager" \
      vendor="loxilb.io" \
      version=$GIT_VERSION \
      release="0.1" \
      summary="loxilb-ingress-manager docker image" \
      description="simple ingress implementation for loxilb" \
      maintainer="backguyn@netlox.io"

WORKDIR /bin/
COPY ./bin/loxilb-ingress-manager /bin/loxilb-ingress-manager
USER root
RUN chmod +x /bin/loxilb-ingress-manager
