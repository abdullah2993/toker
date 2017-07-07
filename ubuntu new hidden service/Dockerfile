FROM abdullah2993/tor:ubuntu
RUN mkdir -p /tor/service/
RUN chmod -R 700 /tor/service/
RUN chown -R debian-tor:debian-tor /tor/service/
RUN echo "HiddenServiceDir /tor/service/" >> /etc/tor/torrc
RUN echo "HiddenServicePort 80 127.0.0.1:65000" >> etc/tor/torrc
COPY helloworld /
ENTRYPOINT service tor restart && while [ ! -f /tor/service/hostname ]; do sleep 1; done && cat /tor/service/hostname && ./helloworld