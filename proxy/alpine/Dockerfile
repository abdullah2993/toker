FROM alpine
LABEL description="Run tor in a container so that client can connect using SOCKS"
LABEL maintainer="a.saleem2993@gmail.com"
RUN apk update
RUN apk add tor
RUN echo "SOCKSPort 0.0.0.0:9050" > /etc/tor/torrc
EXPOSE 9050
CMD ["tor"]