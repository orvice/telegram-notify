FROM orvice/ubuntu-base

RUN mkdir /app/bin -p

WORKDIR /app/bin

ADD notify.sh .
