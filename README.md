### NAT

#### 编译并打包

```sh
# web静态资源
$ cd statics
$ npm install
$ npm run build
# 编译并打包
# 编辑 nat.json 按需填写配置
$ make package
```

最终产物为 `nat.tar.gz`

### 客户端

```sh
# 每隔 10min 发送一次请求
$ while true; do curl -X POST --url "http://server:8787/api/v1/ip_history/add" --data '{"message": "macos"}';sleep 600; done
```

访问 http://117.50.186.230:8787 可以看到你的外网 IP 地址。
