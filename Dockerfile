FROM scratch

MAINTAINER Mitch Anderson <mitch@nuvi.com>

ADD web /

ENV PORT=80
EXPOSE 80
CMD ["/web"]
