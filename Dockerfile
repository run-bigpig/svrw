# 使用 Go 官方镜像作为基础镜像
FROM golang:1.17-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制项目代码到容器中
COPY . .

# 编译 Go 项目
RUN go build -o svrw cmd/main.go

# 使用 alpine 镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 复制编译后的可执行文件到最终镜像中
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./svrw","-addr",":8080"]
