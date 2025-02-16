FROM alpine

ENV TZ Asia/Shanghai
WORKDIR /app

COPY ./go-admin ./go-admin
COPY ./config/settings.yml ./config/settings.yml
COPY ./static ./static
COPY ./temp ./temp

EXPOSE 8000
RUN  chmod 777 go-admin
CMD ["sh","go-admin","server","-c", "config/settings.yml"]