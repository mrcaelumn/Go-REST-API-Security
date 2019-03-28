FROM alpine:3.7

RUN apk --no-cache --update add ca-certificates

ADD Go-REST-API-Security ./
RUN mkdir -p api/swagger
ADD api/swagger/swagger.json api/swagger/swagger.json

ENV GORESTSECURITY_SOCKET tcp://0.0.0.0:8088


EXPOSE 8080

ENTRYPOINT [ "/Go-REST-API-Security", "run" ]
