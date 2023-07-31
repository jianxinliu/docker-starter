FROM alpine:3.16

ENV NAME=hello AGE=11

WORKDIR /app

COPY /app.sh .

RUN chmod +x app.sh

ENTRYPOINT [ "/app/app.sh" ]