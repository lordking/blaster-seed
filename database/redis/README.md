样例--Redis
===========

实现了Redis的set、get、del、publish/subscrbe的简单例子

## 运行前准备

* 需安装redis。
* `config.yaml`，配置文件，配置数据库、日志。

## 依赖库

```shell
$ go get -u github.com/Sirupsen/logrus
$ go get -u github.com/spf13/viper
$ go get -u github.com/garyburd/redigo/redis
```

## 运行样例

```
go run main.go person.go
```
