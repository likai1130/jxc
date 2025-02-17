FROM alpine

# 安装 tzdata 和字体
RUN apk add --no-cache tzdata fontconfig ttf-dejavu
ENV TZ Asia/Shanghai
ENV LANG zh_CN.UTF-8
ENV LANGUAGE zh_CN:zh
ENV LC_ALL zh_CN.UTF-8

WORKDIR /app

# 复制文件
COPY ./go-admin ./go-admin
COPY ./config/settings.yml ./config/settings.yml
COPY ./static ./static
COPY ./temp ./temp

# 设置权限
RUN chmod +x ./go-admin  # 确保可执行权限

# 暴露端口
EXPOSE 8000

# 使用 sh 执行
CMD ["./go-admin", "server", "-c", "config/settings.yml"]