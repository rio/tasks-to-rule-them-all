FROM alpine

RUN apk add bash

COPY bin/echo-server /echo-server

CMD [ "/echo-server" ]
