FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest


WORKDIR /app
# 需要文件COPY到容器内
COPY --from=builder /build/server   ./
COPY --from=builder /build/resource ./resource/
COPY --from=builder /build/configs  ./configs/

# 映射端口
EXPOSE 8888

CMD ["./server"]