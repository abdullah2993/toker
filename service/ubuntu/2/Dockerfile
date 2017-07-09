FROM abdullah2993/tor:ubuntu
RUN mkdir -p /tor/service/
COPY hostname /tor/service/
COPY private_key /tor/service/
RUN chmod -R 700 /tor/service/
RUN chown -R debian-tor:debian-tor /tor/service/
RUN echo "HiddenServiceDir /tor/service/" >> /etc/tor/torrc
RUN echo "HiddenServicePort 80 127.0.0.1:65000" >> etc/tor/torrc
COPY app /
ENTRYPOINT service tor restart && ./helloworld