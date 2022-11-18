##### 主要功能

- 用户注册登录(JWT-Go鉴权)
- 用户基本信息修改，解绑定邮箱，修改密码
- 商品的发布，浏览等
- 购物车的加入，删除，浏览等
- 订单的创建，删除，支付等
- 地址的增加，删除，修改等
- 各个商品的浏览次数，以及部分种类商品的排行
- 设置了支付密码，对用户的金额进行了对称加密
- 支持事务，支付过程发送错误进行回退处理



项目结构



```
gin-mall/
├─api
│  └─v1
├─cache
├─cmd
├─conf
├─dao
├─logs
│  └─2022-11-18.log
├─middleware
├─model
├─pkg
│  ├─e
│  └─util
├─routes
├─serializer
├─service
└─static
    └─imgs
        └─avatar
            └─user1

```





- api : 用于定义接口函数
- cache : 放置redis缓存
- conf : 用于存储配置文件
- cmd : 存放main.go
- dao : 对持久层进行操作
- logs : 存放日志文件
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/e : 封装错误码
- pkg/util : 工具函数
- routes : 路由逻辑处理
- serializer : 将数据序列化为 json 的函数
- service : 接口函数的实现
- static : 存放图片