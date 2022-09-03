# 项目文件说明

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [项目文件说明](#项目文件说明)
    - [sql](#sql)
        - [create_table.sql](#create_tablesql)
            - [创建数据库、创建表结构](#创建数据库创建表结构)
            - [表结构说明](#表结构说明)
            - [注意](#注意)
    - [internal](#internal)
        - [model](#model)
            - [column](#column)
        - [entity](#entity)
        - [service](#service)
            - [context.go](#contextgo)

<!-- markdown-toc end -->

## sql
### create_table.sql
#### 创建数据库、创建表结构

创建数据库的语句：

```sql
CREATE DATABASE `gcbbs`;
```

进入新创建的数据库：

```sql
USE `gcbbs`;
```

然后运行 `sql/create_table.sql` 中的 sql 语句：

```sql
SOURCE /path/to/create_table.sql
```

#### 表结构说明

每张表的介绍：

| 表                    | 简介                                           |
|-----------------------|------------------------------------------------|
| `checkins`            | 签到                                           |
| `comments`            | 评论                                           |
| `follows`             | 关注                                           |
| `integral_logs`       | 积分                                           |
| `likes`               | 给喜欢的资源（文章、评论）添加爱心             |
| `nodes`               | 节点，比如游戏、足球、编程                     |
| `reminds`             | 消息提醒，比如有人点赞了你的评论               |
| `reports`             | 举报                                           |
| `system_notices`      | 系统消息，比如社区发布了新功能                 |
| `system_user_notices` | 针对用户的系统消息提醒，比如你的文章被举报下架 |
| `topics`              | 话题，文章                                     |
| `users`               | 用户                                           |

#### 注意

`users` 表的 `is_admin` 列用于判断该用户是否为社区的管理员，如果是管理员，那么就可以进入 `ip:port/backend`。

## internal
### model

model 包中每个结构体对应数据库中的每张表，为了方便后续使用，我们像这样把它们封装起来：

```go
type userModel struct {
    M     *gorm.DB
    Table string
}
```

- `M` 指向一个 `Users{}`；
- `Table` 是数据库中的表名；

使用一个函数来创建这样的 `userModel` 结构体：

```go
func User() *userModel {
	return &userModel{
		M:     db.DB.Model(&Users{}),
		Table: "users",
	}
}
```

#### column

在 `/internal/model/topic.go` 中，结构体 `Topics` 的字段 `Tags` 的类型是 `column.SA`，在 `/internal/model/column/column.go` 中可以看到，实际上 `column.SA` 类型就是 `[]string` 类型的别名，并且添加了两个方法：`Value`、`Scan`。

### entity
前后端交互的过程中，请求的数据结构不可能只是 model 的结构，还需要其他消息，所以创建实体包。
### service
#### context.go

构建了 `BaseContext` 结构体：

```go
// BaseContext encapsulated context
type BaseContext struct {
	Ctx     *gin.Context
	session sessions.Session
	path    string
}
```

- 使用 `Context` 函数创建一个 `BaseContext` 结构体；

为这个结构体添加了很多方法：

- `Redirect`: 重定向；
- `clear`: 清理session；
- `Back`: 返回上一个url；
- `To`: url跳转；
- `WithError`: 错误消息跳转；
- `WithMsg`: 提示消息跳转；
- `SetAuth`: 设置授权；
- `Auth`: 获取授权；
- `Check`: 检查授权；
- `Forget`: 清理授权；
- `unread`: 是否有消息未读；
- `View`: 模板返回；
- `Json`: 通用的json响应；
- `MDFileJson`: markdown上传图片响应；
