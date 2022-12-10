## Monica 服务治理平台

### 服务注册/发现

1.服务注册

```shell
curl --location --request POST 'http://127.0.0.1:8081/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
    "token": "18ee7064-3cdd-4ed5-a139-fd8d9add5847",
    "namespace": "Test",
    "service_name": "go_wallet_manage_svr",
    "node": {
      "ip": "127.0.0.1",
      "port": 1024,
      "weight": 100,
      "metadata": "[]"
    }
}'
```

2.服务实例更新

```shell
curl --location --request POST 'http://127.0.0.1:8081/api/v1/renew' \
--header 'Content-Type: application/json' \
--data '{
    "token": "18ee7064-3cdd-4ed5-a139-fd8d9add5847",
    "namespace": "Test",
    "service_name": "go_wallet_manage_svr",
    "ip": "127.0.0.1"
}'
```

3.服务实例注销

```shell
curl --location --request POST 'http://127.0.0.1:8081/api/v1/cancel' \
--header 'Content-Type: application/json' \
--data '{
    "token": "18ee7064-3cdd-4ed5-a139-fd8d9add5847",
    "namespace": "Test",
    "service_name": "go_wallet_manage_svr",
    "ip": "127.0.0.1"
}'
```

4.获取服务实例

```shell
curl --location --request GET 'http://127.0.0.1:8081/api/v1/fetch?ns=Test&sname=go_wallet_manage_svr'
```

Response:

```json
{
    "retcode": 0,
    "errmsg": "",
    "data": [
        {
            "ip": "127.0.0.1",
            "port": 1024,
            "weight": 100,
            "metadata": "[]"
        }
    ]
}
```

5.获取服务实例（长轮询）

```shell
curl --location --request GET 'http://127.0.0.1:8081/api/v1/poll?ns=Test&sname=go_wallet_manage_svr'
```

Response: 

```json
{
    "retcode": 0,
    "errmsg": "",
    "data": [
        {
            "ip": "127.0.0.1",
            "port": 1024,
            "weight": 100,
            "metadata": "[]"
        }
    ]
}
```

### 服务部署

1.编译 admin-ui

```shell
yarn install
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

