FROM alpine

# ENV GOPROXY https://goproxy.cn/

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

COPY ./go-admin /go-admin
COPY ./config/settings.yml /config/settings.yml
COPY ./go-admin-db.db /go-admin-db.db
COPY ./static /static
COPY ./temp /temp

EXPOSE 8000
RUN  chmod +x /go-admin
CMD ["/go-admin","server","-c", "/config/settings.yml"]