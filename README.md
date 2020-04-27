# my-gin
自己封装的gin框架

- orm
- mysql
- 引入websocket


# 控制器执行方式

1. 路由找到控制器
2. 控制器执行Init方法，这个不用管，主要作用是NewController反射创建具体的新的控制器对象，将context赋值用的
3. 然后执行，Prepare方法，这个方法可以做一些操作,返回MyException，如果是nil则在程序继续执行，如果非nil,响应错误,写会错误码code
4. 根据路由的:action去执行对应的方法,然后执行完毕
5. 如果:action的方法不存在,会根据请求类型，去执行对应的方法，比如 get请求会执行Get方法,post请求执行Post方法，resetful



# demo运行

```cassandraql
git clone https://github.com/FlashFeiFei/my-gin.git


执行 go run main.go

访问聊天室路由  http://127.0.0.1:8083/ws/home
```