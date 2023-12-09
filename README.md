# billadm

#### 介绍
基于Go语言开发的命令行账单管理工具。一个使用go语言及相关库的练手项目

#### proto定义文件编译命令
```shell
protoc -I proto/ --go_opt=paths=source_relative --go_out=pkg/api/service --go-grpc_out=pkg/api/service proto/storage_service.proto
```

#### v1.0 - 2023/10/02
1.0版本基于命令行操作。数据存储较为简陋，支持功能有限。只能满足基本的记账能力