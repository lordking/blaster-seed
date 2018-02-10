样例--SQLite
============

使用SQLite数据库实现的简单的增、删除、修改、查询例子。

## 运行前的准备

### 编译前安装数据库驱动

开发时，需要提前安装sqlite驱动。编译成功后使用时，无需安装驱动。

```
# Mac
brew install sqlite3
```

### 数据库文件

本例`person.db`已经创建完毕，可以直接使用。如果需要创建，按此下SQL脚本。

```sql
CREATE TABLE "person" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "phone" TEXT
)
```

### 配置文件

`config.yaml`，配置文件，配置数据库、日志。

## 依赖库

```shell
$ go get -u github.com/Sirupsen/logrus
$ go get -u github.com/spf13/viper
$ go get -u github.com/mattn/go-sqlite3
```

## 运行样例

```shell
$ go run main.go person.go
```

## 编译

```shell
# Mac下编译
$ go build --tags "libsqlite3 darwin"
```
