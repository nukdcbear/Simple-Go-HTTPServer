FROM alpine:3.16
LABEL maintainer="DC Barringer <nukdcbear@gmail.com>"

ARG EXPOSED_PORT

WORKDIR /httpserver
COPY blddir/httpserver ./
COPY config.yaml ./

EXPOSE ${EXPOSED_PORT}

ENTRYPOINT [ "/httpserver/httpserver" ]