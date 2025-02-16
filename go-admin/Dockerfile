FROM alpine

ENV TZ Asia/Shanghai
WORKDIR /app

COPY ./go-admin /app/go-admin
COPY ./config/settings.yml /app/config/settings.yml
COPY ./static /app/static
COPY ./temp /app/temp

EXPOSE 8000
RUN  chmod +x /app/go-admin
CMD ["/app/go-admin","server","-c", "/config/settings.yml"]