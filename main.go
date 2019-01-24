package main

import "admin/route"

func main() {

	//g := gin.Default()
	//
	//// 加载静态文件
	//g.Static("/static", "./resources/static")
	//// 加载模板文件
	//g.LoadHTMLGlob("resources/views/**/**/*")
	//
	//g.GET("/", func(c *gin.Context) {
	//
	//	c.HTML(http.StatusOK, "show/show", nil)
	//})
	//
	//g.GET("/test", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "test/test", nil)
	//})
	//
	//err := g.Run(":8080")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	route.Run()

}
