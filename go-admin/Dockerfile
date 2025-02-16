FROM alpine

# 安装 tzdata
RUN apk add --no-cache tzdata

ENV TZ Asia/Shanghai
WORKDIR /app

# 复制文件
COPY ./go-admin ./go-admin
COPY ./config/settings.yml ./config/settings.yml
COPY ./static ./static
COPY ./temp ./temp

# 设置权限
RUN chmod 777 ./go-admin

# 暴露端口
EXPOSE 8000

# 直接执行可执行文件
CMD ["./go-admin", "server", "-c", "config/settings.yml"]