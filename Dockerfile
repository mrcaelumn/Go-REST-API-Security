FROM alpine:3.7

RUN apk --no-cache --update add ca-certificates

ADD go-rest-api-security ./
RUN mkdir -p api/swagger
RUN mkdir file
ADD api/swagger/swagger.json api/swagger/swagger.json
ADD file ./

ENV GORESTSECURITY_SOCKET tcp://0.0.0.0:8088


EXPOSE 8080

ENTRYPOINT [ "/go-rest-api-security", "run" ]
