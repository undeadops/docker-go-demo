FROM scratch

MAINTAINER Mitch Anderson <mitch@nuvi.com>

ADD web /

ENV PORT=8080

CMD ["/web"]
