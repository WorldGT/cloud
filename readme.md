    轻量级云盘系统，基于go-zero、xorm实现。
####命令
    创建API服务
    goctl api new core
    启动服务
    go run core.go -f etc/core-api.yaml
    使用api文件生成代码
    goctl api go -api core.api -dir . -style go_zero

####用户模块
    密码登录
    刷新Authorization
    邮箱注册
    用户详情
    用户容量
####存储池模块
    中心存储池资源管理
    文件上传
   
####个人存储池资源管理
    文件关联存储
    文件列表
    文件名称修改
    文件夹创建
    文件删除
    文件移动
####文件分享模块
    创建分享记录
    获取资源详情
    资源保存