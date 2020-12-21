- 仿照 `macro` 大佬的 `mall` 商城.用Go重写api.不知道能不能实现.干就完事了.

# 项目地址

[项目地址](https://github.com/macrozheng/mall)

# 原接口文档

[文档地址](http://admin-api.macrozheng.com/swagger-ui.html#/)

# 已经实现的接口

- 后台用户管理
    - [x] POST /api/admin/login
    - [x] POST /api/admin/register
    - [x] GET /api/admin/:id
    - [x] POST /api/admin/delete/:id

- 商品管理
    - [x] POST /api/product/create
    - [x] POST /api/product/update/:id
    - [x] GET /api/product/list
    - [x] GET /api/product/simpleList
    - [x] GET /api/product/batchUpdate/deleteStatus -> /api/product/update/deleteStatus 

