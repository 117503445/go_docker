# go_docker

Go 的 Docker 化实践

## 项目地址

<https://github.com/117503445/go_docker>

## 参考博文

<https://www.cnblogs.com/likeli/p/10521941.html>

## 目标

在本地完成调试以后，将代码 Commit 到仓库，然后就会自动 build docker images，然后在生产服务器上自动部署。同时，还要注意配置文件的安全问题。

## 流程

提交代码 -> Docker Hub 的 Automated 服务发现了 Github 上的提交，根据 Dockerfile 构建镜像 -> 生产服务器上的 WatchTower 检测到 Docker Hub 发生更新，自动更新本地的镜像。

## 配置文件传递思路

因为 Github 仓库和 Docker Hub 镜像 都是公开的，所以不可以在这些地方储存配置文件。所以，配置如果通过 docker run 时传递，就可以确保安全性。docker run 时通过 -e 把参数传入 docker 镜像，然后通过 Dockerfile 的定义把参数以命令行的形式传入 Go 程序，Go 程序通过 pflag 接收命令行参数。

## 使用方法

### 生产服务器

使用下列代码运行镜像

```sh
docker rm go_docker -f
docker rmi 117503445/go_docker
docker run -it --name go_docker -d -e var1="dockervar1" -e var2="dockervar1" -p 80:80 --restart=always 117503445/go_docker:latest
```

再配置 WatchTower 以启用自动更新 (以下代码会自动更新所有 docker image)

```sh
docker run -d \
    --name watchtower \
    -v /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower
```

### 本地调试

把 config.yml.example 重命名为 config.yml， 在 config.yml 文件中配置，再按照常规操作运行

使用了分层构建，在 build 层 通过 go build 构筑了 可执行文件 app，再在 prod 层 进行运行。如果以后需要修改配置文件的结构，也需要修改 Dockerfile。

### 本地 Docker 运行

```sh
docker rm go_docker -f
docker rmi 117503445/go_docker

docker build -t 117503445/go_docker . # 国外
docker build -t 117503445/go_docker -f Dockerfile_cn . #国内,启用 go 镜像

docker run -it --name go_docker -d -e var1="dockervar1" -e var2="dockervar1" -p 80:80 --restart=always 117503445/go_docker:latest
```
