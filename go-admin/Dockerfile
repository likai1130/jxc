FROM alpine

ENV TZ Asia/Shanghai

COPY ./go-admin /go-admin
COPY ./config/settings.yml /config/settings.yml
COPY ./static /static
COPY ./temp /temp

EXPOSE 8000
RUN  chmod +x /go-admin
CMD ["/go-admin","server","-c", "/config/settings.yml"]