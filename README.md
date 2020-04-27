# my-gin
自己封装的gin框架

- orm
- mysql
- 引入websocket


# 控制器执行方式

1. 路由找到控制器
2. 控制器执行Init方法，这个不用管，主要作用是NewController反射创建具体的新的控制器对象，将context赋值用的
3. 然后执行，Prepare方法，这个方法可以做一些操作
