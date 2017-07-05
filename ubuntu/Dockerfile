FROM ubuntu:xenial
RUN touch /etc/apt/sources.list.d/tor.list
RUN echo "deb http://deb.torproject.org/torproject.org xenial main" >> /etc/apt/sources.list.d/tor.list
RUN echo "deb-src http://deb.torproject.org/torproject.org xenial main" >> /etc/apt/sources.list.d/tor.list
RUN gpg --keyserver keys.gnupg.net --recv A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89
RUN gpg --export A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89 | apt-key add -
RUN apt-get update
RUN apt-get -y install tor deb.torproject.org-keyring
RUN echo "SOCKSPort 0.0.0.0:9050" > /etc/tor/torrc
EXPOSE 9050
CMD ["tor"]
