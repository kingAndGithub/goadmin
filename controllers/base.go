package controllers

import (
	"admin/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Controller struct{}

type Admin struct{ Controller }
type AdminUser struct{ Controller }

func (s *Controller) display(ctx *gin.Context, args ...interface{}) {

	var m = make(gin.H)
	var data = make(gin.H)
	var name string

	//获取path
	paths := strings.Split(ctx.Request.URL.Path, "/")
	pathsLen := len(paths)

	if pathsLen == 2 {
		if paths[1] == "" {
			paths[1] = "Index"
		}
		paths = append(paths, "index")
	}

	name = paths[1] + "/" + paths[2]

	// 默认渲染自动菜单栏目
	m["_column"] = GetColumn()
	m["_menuSignActive"] = paths[1]
	m["_menuActionSignActive"] = paths[2]

	// 自动识别传参
	argsLen := len(args)
	if argsLen > 0 {

		switch args[0].(type) {
		case string:
			name = args[0].(string)
			if argsLen > 1 {
				data = args[1].(gin.H)
			}
			break
		case gin.H:
			data = args[0].(gin.H)
			if argsLen > 1 {
				name = args[1].(string)
			}
			break
		default:
			break
		}
	}

	for k, v := range data {
		m[k] = v
	}

	ctx.HTML(http.StatusOK, name, m)
}

func GetColumn() auth.Column {
	return auth.Column{
		Name: "首页",
		Menus: []auth.Menu{
			{
				Name:        "管理员",
				HasNode:     true,
				DefaultNode: 1,
				Nodes: []auth.ChildMenu{
					{
						Name:        "管理员列表",
						Sign:        "AdminUser",
						Controller:  AdminUser{},
						DefaultSign: "Index",
						Actions: []auth.Action{
							{
								Name:   "管理员列表",
								Sign:   "Index",
								Method: "GET",
							},
							{
								"管理员详情",
								"Show",
								"GET",
							},
						},
					},
					//{
					//	Name:        "管理员分组",
					//	Sign:        "AdminUserGroup",
					//	Controller:  AdminUser{},
					//	DefaultSign: "Index",
					//	Actions: []auth.Action{
					//		{
					//			Name:   "管理员分组列表",
					//			Sign:   "Index",
					//			Method: "GET",
					//		},
					//	},
					//},
				},
			},
		},
	}
}
