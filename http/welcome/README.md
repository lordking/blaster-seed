样例--welcome
============

一个restapi的样例。

## 1 编译运行前准备

## 1.1 配置说明

所有配置文件均放置在config目录下，内容以YAML格式存放。

```
  +- config
      |
      +---- config.yaml
```

###  配置文件HTTP说明

参数     | 说明
------- | ------------------
http    | HTTP端口。
sslport | HTTPS端口，不能与HTTP相同。
sslcert | HTTPS需要的证书文件的相对路径。
sslkey  | HTTPS需要的公钥文件的相对路径。

ssl_cert和ssl_key的生成方式是：

```
$ go run$GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
```

### 依赖库

```shell
$ go get -u github.com/Sirupsen/logrus
$ go get -u github.com/spf13/viper
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/spf13/cobra

# 单元测试使用
$ go get -u github.com/stretchr/testify/assert
```

## 2 运行样例

编译
```
$ go build
$ ./welcome serve
```

直接运行
```
$ go run main.go serve
```

## 3 单元测试

#### 3.1 单元测试

```
$ cd test
$ go test -v -test.run Test_Hello
```

#### 3.2 性能测试

```
$ cd benchmark/benchmark
$ go run main.go -m RequestHello
```