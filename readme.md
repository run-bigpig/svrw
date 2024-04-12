# SVRW

该项目是一个基于Go开发,获取短视频无平台水印的demo。

## 安装

首先，您需要安装 Go 编程语言和相关工具。您可以从 Go 的官方网站 [https://golang.org/](https://golang.org/) 下载并安装 Go。

在安装 Go 后，您需要设置好 Go 的工作目录和环境变量。请参考 Go 官方文档中的说明 [https://golang.org/doc/install](https://golang.org/doc/install)。

安装完成后，您可以通过以下命令获取该项目：
```
git clone https://github.com/run-bigpig/svrw.git
```

## 编译

进入项目的根目录，执行以下命令编译项目：
```
go build -o svrw cmd/main.go
```

编译完成后，将在项目根目录生成可执行文件。

## 使用

### 运行项目

执行以下命令来运行项目：
```
./svrw
```
### 可选参数
```
-addr string
    监听地址 (default ":10800")
-bc string
    哔哩哔哩的cookie (default "")
```
然后，您可以在浏览器中访问 `http://localhost:10800`（默认端口）来查看项目运行情况。

### 获取短视频无平台水印
```
http://localhost:8080/api?url=https://v.douyin.com/iYdNyd34/
```

### 目前支持的平台
- [x] 抖音
- [x] 快手
- [x] 皮皮虾
- [x] 微视
- [x] 最右
- [x] 西瓜视频(封面地址有时效性、视频地址访问需要带referer:https://www.ixigua.com/)
- [x] bilibili 哔哩哔哩
- [x] 小红书(解析的地址有些需要下载之后才可以改后缀播放)
- [ ] 陆续完善中...

## 贡献

如果您对该项目有任何建议或发现了 Bug，欢迎提交 Issue 或 Pull Request。

## 免责声明

本仓库只为学习研究，如涉及侵犯个人或者团体利益，请与我取得联系，我将主动删除一切相关资料，谢谢！

