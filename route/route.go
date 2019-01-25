package route

import (
	"admin/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

var controllerSet = make(map[string]interface{})
var methods = make(map[string]map[string]reflect.Method)

func Run() {

	g := gin.Default()

	// 加载静态文件
	g.Static("/static", "./resources/static")
	// 加载模板文件
	g.LoadHTMLGlob("resources/views/**/**/*")

	column := controllers.GetColumn()
	for _, menus := range column.Menus {
		for _, childs := range menus.Nodes {
			controllerSet[childs.Sign] = childs.Controller
			methods[childs.Sign] = make(map[string]reflect.Method)
			for _, action := range childs.Actions {

				m := reflect.TypeOf(childs.Controller)
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

	g.GET("/", controllers.AdminUser{}.Index)

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
	pathsLen := len(paths)

	if pathsLen == 2 {
		if paths[1] == "" {
			paths[1] = "Index"
		}
		paths = append(paths, "index")
	}

	//反射进行调用
	methods[paths[1]][paths[2]].Func.Call([]reflect.Value{
		reflect.ValueOf(controllerSet[paths[1]]), // 映射的结构体
		reflect.ValueOf(ctx), // 传参
	})
}
