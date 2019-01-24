package auth

// 栏目
type Column struct {
	Name  string // 名称
	Menus []Menu // 所有菜单
}

// 菜单栏目
type Menu struct {
	Name        string      // 名称
	DefaultNode int         // 默认栏目
	Nodes       []ChildMenu // 所有子栏目
}

// 菜单子栏目
type ChildMenu struct {
	Name     string   // 名称
	Sign     string   // 路由标识位
	Actions []Action // 所有操作
}

// 单项操作
type Action struct {
	Name string // 名称
	Sign string // 路由标识位
	Method string // 操作方式 POST/GET
}

func Init() Column {

	return Column{"首页", []Menu{
		{
			"管理员",
			1,
			[]ChildMenu{
				{
					"管理员列表",
					"AdminUser",
					[]Action{
						{
							"管理员列表",
							"Index",
							"GET",
						},
						{
							"管理员详情",
							"Show",
							"GET",
						},
						{
							"管理员添加",
							"Create",
							"POST",
						},
						{
							"管理员更新",
							"Update",
							"POST",
						},
						{
							"管理员删除",
							"Del",
							"POST",
						},
					},
				},
			},
		},
	}}
}
