package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	"github.com/nats-io/nats.go"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.2
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	logger.Setup()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//global.GVA_DB = global.GVA_DB.Debug()
	connect, err := nats.Connect(global.GVA_CONFIG.Nats.Url)
	if err != nil {
		panic(err)
	}
	global.NatsClient = connect
	defer connect.Close()
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		go func() {
			initialize.RegisterTables() // 初始化表
		}()
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	defer cache.ProxyCache.Clear()
	core.RunWindowsServer()
	//select {}
}
