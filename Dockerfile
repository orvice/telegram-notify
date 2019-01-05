FROM orvice/ubuntu-base

RUN mkdir /app/bin

WORKDIR /app/bin

ADD notify.sh .
