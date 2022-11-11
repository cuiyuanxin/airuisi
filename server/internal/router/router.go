package router

import (
	"time"

	"github.com/cuiyuanxin/airuisi/global"

	"github.com/cuiyuanxin/airuisi/internal/middleware"

	"github.com/cuiyuanxin/airuisi/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// 测试通信
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})
	// 限流器
	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})

	// 链路追踪
	r.Use(middleware.Tarcing())
	if global.ServerSetting.GinMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout * time.Second))
	r.Use(middleware.Cors())

	// 后台
	//login := v1.NewLogin()
	//sysUser := v1.NewSysUser()
	//sysRole := v1.NewSysRole()
	//sysMenu := v1.NewSysMenu()
	//sysApi := v1.NewSysApi()
	//natPos := v1.NewNntPos()
	//population := v1.NewPopulation()
	//workplace := v1.NewWorkplace()
	//statistical := v1.NewStatistical()
	//village := v1.NewVillage()
	//
	//apiR := r.Group("/api")
	//// 获取验证码
	//apiR.POST("/sms", login.SendSms)
	//// 登录
	//apiR.POST("/login", login.Login)
	//
	//apiRv1 := apiR.Group("/v1")
	//apiRv1.Use(middleware.JWT())
	//
	//apiRv1.Use(middleware.CasbinAuth())
	//{
	//	// 获取首页数据统计
	//	apiRv1.GET("/workplace/statistical", workplace.Statistical)
	//	// 获取管理员信息
	//	apiRv1.GET("/userinfo", sysUser.UserInfo)
	//	// 获取nav
	//	apiRv1.GET("/nav", sysMenu.Nav)
	//
	//	//管理员分组
	//	// 获取管理员列表
	//	apiRv1.GET("/users", sysUser.List)
	//	// 添加管理员
	//	apiRv1.POST("/users", sysUser.Create)
	//	// 获取修改管理员信息
	//	apiRv1.GET("/users/:id", sysUser.Info)
	//	// 修改管理员信息
	//	apiRv1.PUT("/users/:id", sysUser.Update)
	//	// 删除管理员
	//	apiRv1.DELETE("/users", sysUser.Delete)
	//
	//	// 角色组分组
	//	// 获取角色组列表
	//	apiRv1.GET("/role", sysRole.List)
	//	apiRv1.GET("/role/list", sysRole.Lists)
	//	// 添加角色组
	//	apiRv1.POST("/role", sysRole.Create)
	//	// 编辑角色
	//	apiRv1.GET("/role/:id", sysRole.Info)
	//	apiRv1.PUT("/role/:id", sysRole.Update)
	//	// 删除角色
	//	apiRv1.DELETE("/role", sysRole.Delete)
	//	// 修改角色状态
	//	apiRv1.POST("/role/status", sysRole.State)
	//	// 授权
	//	apiRv1.POST("/role/permissions", sysRole.Permissions)
	//
	//	// 菜单
	//	// 获取菜单列表
	//	apiRv1.GET("/menu", sysMenu.List)
	//	apiRv1.GET("/menu/tree", sysMenu.ListTree)
	//	// 添加菜单
	//	apiRv1.POST("/menu", sysMenu.Create)
	//	// 编辑菜单
	//	apiRv1.GET("/menu/:id", sysMenu.Info)
	//	apiRv1.PUT("/menu/:id", sysMenu.Update)
	//	// 删除菜单
	//	apiRv1.DELETE("/menu", sysMenu.Delete)
	//
	//	// Api
	//	// 获取接口列表
	//	apiRv1.GET("/api", sysApi.List)
	//	apiRv1.GET("/apilist", sysApi.Lists)
	//	// 添加接口
	//	apiRv1.POST("/api", sysApi.Create)
	//	// 编辑接口
	//	apiRv1.GET("/api/:id", sysApi.Info)
	//	apiRv1.PUT("/api/:id", sysApi.Update)
	//	// 删除接口
	//	apiRv1.DELETE("/api", sysApi.Delete)
	//
	//	//检查点
	//	// 获取检查点列表
	//	apiRv1.GET("/natpos", natPos.List)
	//	apiRv1.GET("/natpos/list", natPos.Lists)
	//	// 添加检查点
	//	apiRv1.POST("/natpos", natPos.Create)
	//	// 获取修改检查点信息
	//	apiRv1.GET("/natpos/:id", natPos.Info)
	//	// 修改检查点信息
	//	apiRv1.PUT("/natpos/:id", natPos.Update)
	//	// 删除检查点
	//	apiRv1.DELETE("/natpos", natPos.Delete)
	//
	//	//人口管理
	//	// 获取人员管理列表
	//	apiRv1.GET("/personnel", population.List)
	//	// 异步导出
	//	apiRv1.POST("/personnel/export", population.ExportCsv)
	//	apiRv1.POST("/personnel/nodify", population.GetNodifyData)
	//	apiRv1.POST("/personnel/download", population.Download)
	//	apiRv1.POST("/personnel/finishDownload", population.FinishDownload)
	//	// 同步导出数据
	//	apiRv1.POST("/personnel/export1", population.ExportCsv1)
	//
	//	// 数据统计
	//	apiRv1.GET("/statistical", statistical.List)
	//	apiRv1.POST("/statistical/export", statistical.ExportCsv)
	//
	//	// 列表
	//	apiRv1.GET("/population", population.PopulationList)
	//	// 添加
	//	apiRv1.POST("/population", population.PopulationCreate)
	//	// 获取修改信息
	//	apiRv1.GET("/population/:id", population.PopulationInfo)
	//	// 修改信息
	//	apiRv1.PUT("/population/:id", population.PopulationUpdate)
	//	// 删除
	//	apiRv1.DELETE("/population", population.PopulationDelete)
	//	apiRv1.GET("/population/export", population.ExportPopulationCsv)
	//
	//	// 村/居
	//	apiRv1.GET("/village", village.List)
	//	apiRv1.GET("/village/list", village.Lists)
	//	apiRv1.POST("/village", village.Create)
	//	apiRv1.GET("/village/:id", village.Info)
	//	apiRv1.PUT("/village/:id", village.Update)
	//
	//}

	// vue
	//r.LoadHTMLGlob("template/admin/index.html")
	//r.Static("assets", "template/admin/assets")
	////r.StaticFS("assets", http.Dir("template/admin"))
	//r.GET("admin", func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})

	//r.StaticFS("/static/qrcode", http.Dir(global.WxInfoSetting.QrcodeSavePath))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
