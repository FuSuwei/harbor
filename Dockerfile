FROM ubuntu:latest

ENV TZ Asia/Shanghai
RUN mkdir -p harbor
WORKDIR /harbor

COPY harbor /harbor/.

COPY docs /harbor/docs
COPY statics /harbor/statics

EXPOSE 8000

CMD ["./harbor"]