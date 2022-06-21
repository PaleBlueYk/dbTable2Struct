# dbTable2Struct

## 简介
将数据库中表转换成Golang结构体(目前仅支持Clickhouse,Mysql版本应该能找到其他现成的)    
背景: 开发中用到Clickhouse数据库,已知数据库中有相关表，但没有现成的结构体，一共39张表，字段繁多(通过本项目生成的结果有1200行左右，谁手写谁**)

## 使用方法

### 1. 直接拉取代码运行

1. 拉取项目

```shell
git clone https://github.com/PaleBlueYk/dbTable2Struct.git
```

2. 修改配置文件 配置文件位于项目目录下 conf.toml

3. 运行程序

```shell
go run main.go
```

### 2. 直接下载可执行文件(待更新)

## 实现原理

1. 使用gorm连接数据库，通过`show tables`和`desc <table_name>`来获取所有表名和字段
2. 完善ymodel/obj.go中的Objs结构体来准备需要使用的结构体

```go
type Objs struct {
PkgName string   // 包名
Imp     []string // 导入包
Objs    []St // 结构体
}

type St struct {
ObjName    string // 结构体名
ObjExtFrom string // 结构体继承自
FieldList  []Filed // 字段
}

type Filed struct {
FieldName string // 字段名
FieldType string // 字段类型
FieldTag  string // 字段tag
}
```

3. 使用`go template`来生成文件

## 声明

本项目前主要用于本人的个人项目，因此实现写在example中，如有不同需要可以自己改改