package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/flipped-aurora/gin-vue-admin/server/router/runner"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	r := runner.Runner{}
	{
		biz_apphubRouter := router.RouterGroupApp.Biz_apphub
		biz_apphubRouter.InitBizAppHubRouter(privateGroup, publicGroup)
		biz_apphubRouter.InitBizToolCmdSrvApiRouter(privateGroup, publicGroup)
		biz_apphubRouter.InitBizCloudFunctionRouter(privateGroup, publicGroup)
		r.InitRunnerRouter(privateGroup, publicGroup)
	}
}

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
