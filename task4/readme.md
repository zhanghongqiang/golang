1.创建数据库：CREATE DATABASE blog_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
2.配置文件 configs/config.yaml
3.运行应用 
   mkdir -p logs
   go mod tidy
   go run cmd/main.go 

   10. API 接口文档
4. 用户认证

  注册用户: POST /api/v1/register

  用户登录: POST /api/v1/login

5 文章管理

获取所有文章: GET /api/v1/posts(公开)

获取单篇文章: GET /api/v1/posts/:id(公开)

创建文章: POST /api/v1/posts(需要认证)

更新文章: PUT /api/v1/posts/:id(需要认证，仅作者)

删除文章: DELETE /api/v1/posts/:id(需要认证，仅作者)

6 评论管理

获取文章评论: GET /api/v1/posts/:post_id/comments(公开)

创建评论: POST /api/v1/posts/:post_id/comments(需要认证)

7 错误处理

系统包含以下错误处理：

400: 请求参数错误

401: 认证失败

403: 权限不足

404: 资源不存在

409: 资源冲突

500: 服务器内部错误
