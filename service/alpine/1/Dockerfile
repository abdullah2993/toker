FROM golang:1.8-alpine AS build
RUN mkdir /src
WORKDIR /src
COPY main.go /src
RUN go build -o app

FROM abdullah2993/tor:alpine
RUN mkdir -p /tor/service/
RUN chmod -R 700 /tor/service/
RUN echo "HiddenServiceDir /tor/service/" >> /etc/tor/torrc
RUN echo "HiddenServicePort 80 127.0.0.1:65000" >> etc/tor/torrc
COPY --from=build /src/app /
ENTRYPOINT $(tor > /dev/null &) && while [ ! -f /tor/service/hostname ]; do sleep 1; done && cat /tor/service/hostname && ./app