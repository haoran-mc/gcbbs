### 2022-08-30

- 第一次提交，今天没有开始写代码，把项目的目录结构大致看了看；

### 2022-08-31

- 创建了数据库，新建表结构，表结构的说明可以在[这里找到（project_documents_description）](./proj_docs_desc.md)；

### 2022-09-01

- 完成配置文件 `config/config.yaml`，然后在 `pkg/config/config.go` 中读取配置文件；
- 使用 `pkg/db/mysql.go`、`pkg/redis/redis.go` 中的 init 函数初始化 mysql 和 redis；
- 这个项目使用 go 内置的 `html/template` 渲染 html，我们会在工具包 `pkg/utils` 中写一些函数以便能够在 template 中完成一些复杂的操作；

### 2022-09-02

- `assets`、`views` 前端的部分拿来主义了 `:D`；
- `cmd/main.go`、`cmd/webserver/server.go` 这是整个项目的开始，配置静态文件位置、模板、模板函数、session 中间件，然后开启一个服务；

### 2022-09-03

- `internal/model/*`，所有 model 代码，对应数据库中的每个表项；
- 添加 `/docs/v1.0/code_diary.md`，记录我完成这个项目的过程；
- 添加 `/docs/v1.0/proj_docs_desc.md`，对项目文件说明；

### 2022-09-04

- 添加 `internal/service/context.go`，封装 context 上下文，添加鉴权、处理 session 等方法；
- `internal/entity/frontend`，前台需要用到的实体；
- `internal/consts`，常量；
- `internal/service/frontend/checkin.go`，关于签到的服务层代码；

### 2022-09-05

- 添加 `internal/subject/*`，TODO 这个包是干啥的？
- `internal/service/frontend/comment.go`，关于评论的服务层代码；
