## Monica 服务治理平台

1.编译 admin-ui

```shell
yarn build
mv dist/ ../dist 
```

2.二进制包部署

```shell
go build -o monica
chmod +x restart.sh
./restart.sh
```

3.Docker 部署

```shell
# 编译镜像
docker build -t ccr.ccs.tencentyun.com/tcb-xxx-xupz/prod-xxx-online:[tag] .

# 推送镜像
docker push ccr.ccs.tencentyun.com/tcb-xxx-xupz/prod-xxxx-online:[tag]
```