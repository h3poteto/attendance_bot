FROM h3poteto/golang:1.11.2

USER root

ADD . /go/src/github.com/h3poteto/attendance_bot

WORKDIR /go/src/github.com/h3poteto/attendance_bot

RUN chown -R go:go /go/src/github.com/h3poteto

USER go

RUN set -ex && \
    dep ensure && \
    go build


CMD [ "./attendance_bot" ]
