- 仿照 `macro` 大佬的 `mall` 商城.用Go重写api.不知道能不能实现.干就完事了.

# 项目地址

[项目地址](https://github.com/macrozheng/mall)

# 原接口文档

[文档地址](http://admin-api.macrozheng.com/swagger-ui.html#/)

# 已经实现的接口

- 后台用户管理
    - [x] GET /api/admin/:id -> GET /api/admin/user/:id
    - [x] POST /api/admin/delete/:id
    - [x] GET /api/admin/info (用户菜单和角色还没有获取)
    - [ ] GET /api/admin/list
    - [x] POST /api/admin/login
    - [ ] POST /api/admin/logout
    - [ ] GET /api/admin/refreshToken
    - [x] POST /api/admin/register
    - [ ] GET /api/admin/role/:adminId
    - [ ] POST /api/admin/role/update
    - [ ] POST /api/admin/update/:id
    - [ ] POST /api/admin/updatePassword
    - [ ] POST /api/admin/updateStatus/:id

- 商品管理
    - [x] POST /api/product/create
    - [x] GET /api/product/list
    - [x] GET /api/product/simpleList
    - [x] POST /api/product/update/:id
    - [x] POST /api/product/batchUpdate/deleteStatus -> /api/product/update/deleteStatus
    - [x] POST /api/product/batchUpdate/newStatus -> /api/product/update/newStatus
    - [x] POST /api/product/batchUpdate/publishStatus -> /api/product/update/publishStatus
    - [x] POST /api/product/batchUpdate/recommendStatus -> /api/product/update/recommendStatus
    - [x] POST /api/product/batchUpdate/verifyStatus -> /api/product/update/verifyStatus
    - [x] GET /api/product/updateInfo/:id

