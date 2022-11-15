FROM alpine:3.16
LABEL maintainer="DC Barringer <nukdcbear@gmail.com>"

WORKDIR /httpserver
COPY blddir/httpserver ./
COPY config.yaml ./

EXPOSE 3000

ENTRYPOINT [ "/httpserver/httpserver" ]