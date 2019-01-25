package auth

// 栏目
type Column struct {
	Name  string // 名称
	Menus []Menu // 所有菜单
}

// 菜单栏目
type Menu struct {
	Name        string      // 名称
	HasNode		bool        // 是否有子栏目
	DefaultNode int         // 默认栏目
	Nodes       []ChildMenu // 所有子栏目
}

// 菜单子栏目
type ChildMenu struct {
	Name        string      // 名称
	Sign        string      // 路由标识位
	Controller  interface{} // 控制器
	DefaultSign string      // 默认显示路由
	Actions     []Action    // 所有操作
}

// 单项操作
type Action struct {
	Name   string // 名称
	Sign   string // 路由标识位
	Method string // 操作方式 POST/GET
}
