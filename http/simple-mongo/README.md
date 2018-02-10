样例--simple-mongo
==================

simple-mongo样例，一个最基础的mongodb + http rest api的样例

## 1 编译运行前准备

## 配置说明

所有配置文件均放置在config目录下，内容以YAML格式存放。

```
  +- config
      |
      +---- config.yaml
```

###  `http`配置说明

参数     | 说明
------- | ------------------
http    | HTTP端口。
sslport | HTTPS端口，不能与HTTP相同。
sslcert | HTTPS需要的证书文件的相对路径。
sslkey  | HTTPS需要的公钥文件的相对路径。

ssl_cert和ssl_key的生成方式是：

```
 $ go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
```

### `database`配置说明

参数      | 说明
-------- | ------------------
url      | 数据库URL连接字符串
database | 数据库名

连接字符串的格式是

`[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]`

### 依赖库

```shell
$ go get -u github.com/Sirupsen/logrus
$ go get -u github.com/spf13/viper
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/spf13/cobra
$ go get -u gopkg.in/mgo.v2

# 单元测试使用
$ go get -u github.com/stretchr/testify/assert
```

## 2 编译运行

```shell
$ go build
$ ./simple-mongo serve
```

或

```shell
$ go run main.go serve
```

## 3 测试

### 3.1 单元测试

```shell
 $ cd test
 $ go test -v --test.run Test_Create
 $ go test -v --test.run Test_Find
 $ go test -v --test.run Test_Update
 $ go test -v --test.run Test_Delete
```

### 3.2 性能测试

```shell
 $ cd benchmark/benchmark
 $ go run main.go -m RequestCreate
 $ go run main.go -m RequestFind
 $ go run main.go -m RequestUpdate
 $ go run main.go -m RequestDelete
```
