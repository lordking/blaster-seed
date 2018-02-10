样例--MongoDB
================

使用MongoDB数据库实现的简单的增、删除、修改、查询例子。


## 运行前准备

* 需安装MongoDb数据库。
* `config.yaml`，配置文件，配置数据库、日志。

## 依赖库

```shell
$ go get -u github.com/Sirupsen/logrus
$ go get -u github.com/spf13/viper
$ go get -u gopkg.in/mgo.v2
```

## 运行样例

```
go run main.go person.go
```
