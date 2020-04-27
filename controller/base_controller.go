package controller

import (
	"errors"
	"fmt"
	"github.com/FlashFeiFei/my-gin/exception"
	"github.com/FlashFeiFei/my-gin/help"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//创建一个控制器
//参数是一个执行的controller
func NewController(exec_controller Controller) func(*gin.Context) {

	return func(context *gin.Context) {

		//通过反射创建一个新的controller，为什么要这样做？
		//因为如果不这么做，所有用户都公用一个controller，又因为gin.context是指针，所以用户第二次请求会覆盖第一次请求的context
		exec_controller_type := reflect.TypeOf(exec_controller) //获取controller的指针的reflect.Type
		trueType := exec_controller_type.Elem()                 //获取controller的真实类型
		ptrValue := reflect.New(trueType)                       //获取controller的真实值
		controller := ptrValue.Interface().(Controller)         //底层的“值” =>  转interface{} => 再转具体类型 Controller

		defer func() {
			//捕获异常 try/catch
			if err := recover(); err != nil {
				log.Println(err)                                    //异常日志输出
				help.Gin500ErrorResponse(context, err.(error), nil) //500响应
			}
			//final
			controller.Finish() //做一些释放资源的操作
		}()

		controller.Init(context)               //控制器初始化,类似构造函数
		base_exception := controller.Prepare() //一些钩子吧,在真正执行到控制器请求前在做一下操作，例如权限认证等

		if base_exception != nil {
			help.Gin200ErrorResponse(context, base_exception.GetCode(), base_exception.Error(), nil)
			return //结束请求
		}

		var param []reflect.Value         // 反射调用方法所需要的参数
		action := context.Param("action") //获取执行控制器的方法
		id_string := context.Param("id")
		log.Println(id_string)
		id, err := strconv.ParseUint(id_string, 10, 64) //id
		if err == nil {
			//转成功才添加
			param = make([]reflect.Value, 1)
			param[0] = reflect.ValueOf(id)
		}

		log.Println(id)
		if len(action) == 0 {
			//没有找到action参数，通过请求类型去执行具体对应的方法
			switch method := context.Request.Method; method {
			case http.MethodGet:
				controller.Get()
			case http.MethodPost:
				controller.Post()
			case http.MethodDelete:
				controller.Delete()
			case http.MethodHead:
				controller.Head()
			case http.MethodPatch:
				controller.Patch()
			case http.MethodPut:
				controller.Put()
			case http.MethodOptions:
				controller.Options()
			default:
				panic(errors.New(fmt.Sprintf("还不支持的方法: %s", method)))
			}
			return
		}

		//通过反射去执行具体的方法
		value := reflect.ValueOf(controller)
		//通过路径中的action参数，去解析，调用具体的控制器方法
		action = strings.Trim(action, " ") //修剪一下空格
		var action_split []string
		action_split = strings.Split(action, "-") // '-' 符号分割方法, 例如 hello-world，谷歌是 '-'
		if len(action_split) == 1 {
			action_split = strings.Split(action, "_") // '_'符号分割方法，例如hello_world,百度是 '_'
		}

		for index, item := range action_split {
			action_split[index] = help.StrFirstToUpper(item) //首字母大写
		}
		action = strings.Join(action_split, "") //拼接方法
		call_method := value.MethodByName(action)
		log.Println("执行方法", action)
		if !call_method.IsValid() {
			help.Gin400NotFoundResponse(context, errors.New(fmt.Sprintf("找不到方法:%s", action)), nil)
			return
		}
		//调用方法
		call_method.Call(param)
		return
	}
}

//参考beego
type Controller interface {
	Init(ctx *gin.Context)            //ctx是gin的Context controller是当前执行的控制器,初始化
	Prepare() exception.BaseException //解析
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish() //这个函数是在执行完相应的 HTTP Method 方法之后执行的，默认是空，用户可以在子 struct 中重写这个函数，执行例如数据库关闭，清理数据之类的工作。
}

type BaseController struct {
	Ctx *gin.Context //gin框架的Context
}

//初始化函数
func (c *BaseController) Init(ctx *gin.Context) {
	c.Ctx = ctx
}

//做一些鉴权操作等
func (c *BaseController) Prepare() exception.BaseException {
	return nil
}

func (c *BaseController) Post() {
	help.Gin200SuccessResponse(c.Ctx, "Post请求成功", nil)
	return
}

//默认执行的Get方法
func (c *BaseController) Get() {
	help.Gin200SuccessResponse(c.Ctx, "Get请求成功", nil)
	return
}

func (c *BaseController) Delete() {
	help.Gin200SuccessResponse(c.Ctx, "Delete请求成功", nil)
	return
}

func (c *BaseController) Put() {
	help.Gin200SuccessResponse(c.Ctx, "Put请求成功", nil)
	return
}

func (c *BaseController) Head() {
	help.Gin200SuccessResponse(c.Ctx, "Head请求成功", nil)
	return
}

func (c *BaseController) Patch() {
	help.Gin200SuccessResponse(c.Ctx, "Path请求成功", nil)
	return
}

func (c *BaseController) Options() {
	help.Gin200SuccessResponse(c.Ctx, "Options请求成功", nil)
	return
}

func (c *BaseController) Finish() {
	//可以做一些释放资源的操作
}
