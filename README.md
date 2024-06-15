# Temple-Go-Web-Project(TGWP)

`TGWP`是一个轻量级`GOLang`网络框架脚手架，旨在专注业务代码开发，免去重复逻辑书写

## 快速开始
1. 前置条件
```
go 1.22+
```
2. Fork本仓库到自己指定项目仓库名

3. 修改 `go.mod` 文件中 `module tgwp`为自己的项目名称
4. 使用全局替换 将 `tgwp` 替换为刚刚修改的字段
5. 将`config.yaml.template`更改为`config.yaml`并且填写相关配置
6. 安装相关依赖
```shell
go mod tidy
go mod install
```
7. 编写业务代码

## 目录结构
```
├── README.md
├── cmd       程序入口
├── configs   存放配置实体类
├── db        存放数据库相关
├── global    存放全局变量和常量
├── initalize 程序初始化文件
├── internal  核心业务代码
├── log       log文件配置
├── pkg       第三方通用包
├── test      测试文件
├── utils     工具类
├── config.yaml.template 配置文件
├── go.mod
└── go.sum
```

## 目前集成的功能

1. 数据库链接（使用gorm）
2. Redis链接（使用go-redis/v8）
3. 日志系统（使用zap logrus库分别实现 日志分割 日志轮转）

## 待开发功能
- [ ] hertz gin fiber 相关基建逻辑（@Dbinggo）
- [ ] kafka rocketmq rabbitmq (nobody)
- [ ] 第三方工具类扩展
- [ ] 优化目录结构