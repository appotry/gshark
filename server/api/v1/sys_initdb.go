package v1

import (
	"github.com/madneal/gshark/global"
	"github.com/madneal/gshark/model/request"
	"github.com/madneal/gshark/model/response"
	"github.com/madneal/gshark/service"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("非法访问")
		response.FailWithMessage("非法访问", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := service.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("自动创建数据库失败", zap.Any("err", err))
		err = service.RemoveDB(dbInfo)
		if err != nil {
			global.GVA_LOG.Error("移除数据库", zap.Error(err))
		}
		response.FailWithMessage("自动创建数据库失败，请查看后台日志", c)
		return
	}
	response.OkWithData("自动创建数据库成功", c)
}

func CheckDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Info("数据库无需初始化")
		response.OkWithDetailed(gin.H{
			"needInit": false,
		}, "数据库无需初始化", c)
		return
	} else {
		global.GVA_LOG.Info("前往初始化数据库")
		response.OkWithDetailed(gin.H{
			"needInit": true,
		}, "前往初始化数据库", c)
		return
	}
}
