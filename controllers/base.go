package controllers

type Controller struct{}

type Admin struct{ Controller }
type AdminUser struct{ Controller }

// 控制器
var Controllers = map[string]interface{}{
	"Admin":     Admin{},
	"AdminUser": AdminUser{},
}
