package route

import (
	"admin/auth"
	"admin/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

// 加载所有控制器
var controllerSet = controllers.Controllers

var methods = make(map[string]map[string]reflect.Method)

func Run() {

	g := gin.Default()

	// 加载静态文件
	g.Static("/static", "./resources/static")
	// 加载模板文件
	g.LoadHTMLGlob("resources/views/**/**/*")

	column := auth.Init()
	for _, menus := range column.Menus {
		for _, childs := range menus.Nodes {
			methods[childs.Sign] = make(map[string]reflect.Method)
			for _, action := range childs.Actions {
				m := reflect.TypeOf(controllerSet[childs.Sign])
				mf, exist := m.MethodByName(action.Sign)
				if !exist {
					panic("不存在的方法：" + childs.Sign + "/" + action.Sign)
				}
				methods[childs.Sign][strings.ToLower(action.Sign)] = mf
				switch action.Method {
				case "POST":
					g.POST("/"+childs.Sign+"/"+strings.ToLower(action.Sign), doHandle)
					break
				case "GET":
					g.GET("/"+childs.Sign+"/"+strings.ToLower(action.Sign), doHandle)
					break
				default:
					break
				}
			}
		}
	}

	g.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "show/show", nil)
	})

	g.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test/test", nil)
	})

	err := g.Run(":8080")

	if err != nil {
		fmt.Println(err)
	}
}

func doHandle(ctx *gin.Context) {
	//获取path
	paths := strings.Split(ctx.Request.URL.Path,"/")

	fmt.Println(paths[1], paths[2])

	vals := make([]reflect.Value, 2)

	vals[0] = reflect.ValueOf(controllerSet[paths[1]])
	vals[1] = reflect.ValueOf(ctx)

	fmt.Println(vals[0])
	fmt.Println(methods)
	//反射进行调用
	methods[paths[1]][paths[2]].Func.Call(vals)
}
