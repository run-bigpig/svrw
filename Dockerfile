# 使用 Go 官方镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制项目代码到容器中
COPY . .

#修改go mod 代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 编译 Go 项目
RUN go build -o svrw cmd/main.go

# 暴露端口
EXPOSE 8080
# 启动应用
CMD ["./svrw","-addr",":8080"]
