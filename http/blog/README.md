样例---blog
===========

blog样例，是mongodb + restapi + webapp的样例。

## 1 编译运行前准备

## 1.1  安装js库

运行之前，除go语言环境外先要安装js库。步骤如下：

- 如果没有nodejs，到如下地址下载安装。

  <https://nodejs.org>

- 如果没有bower, 安装。

$ sudo npm install -g bower

- 安装js库

$ cd blog $ bower install

## 1.2 配置说明

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

```bash
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
$ go get -u github.com/gin-gonic/contrib/static

# 单元测试使用
$ go get -u github.com/stretchr/testify/assert
```

## 2 编译运行

```bash
$ go build
$ ./blog serve
```

或者

```
$ go run main.go serve
```

成功后，可访问如下地址:

<http://localhost:8000/login.html>

用户名/密码: admin/admin

## 3 单元测试

### 3.1 用户登录接口测试

```bash
$ cd test
$ go test -v -test.run Test_Login
```

如果单元测试运行成功，将会在终端上的打印输出中获得token。这将用于后面的单元测试。 打开test.go文件，修改token的值。如：

```golang
  token = "57884dba17a06faba180e46a"
```

#### 3.2 创建日志

```
$ go test -v -test.run Test_Create
```

#### 3.3 查询日志

```bash
$ go test -v -test.run Test_Find
```

#### 3.4 修改日志

在运行测试之前，先通过之前创建或者查询的测试用例获取一个id。然后打开test.go文件，修改update_id的值。如：

```golang
  update_id := "57884d1a17a06faba180e468"
```

再运行一下测试

```bash
$ go test -v -test.run Test_Update
```

#### 3.5 删除日志

在运行测试之前，先通过之前创建或者查询的测试用例获取一个id。然后打开test.go文件，找到delete_id的值。如：

```golang
  delete_id = "57884d1a17a06faba180e468"
```

再运行一下测试

```bash
$ go test -v -test.run Test_Delete
```
