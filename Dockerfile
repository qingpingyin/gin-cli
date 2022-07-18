FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y


WORKDIR /app

# 需要文件COPY到容器内
COPY ./configs /app/configs
COPY ./resource /app/resource

# 映射端口
EXPOSE

CMD ["./server"]