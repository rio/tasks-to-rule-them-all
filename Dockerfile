FROM scratch

COPY bin/echo-server /echo-server

CMD [ "/echo-server" ]
