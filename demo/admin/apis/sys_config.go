package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"admin/models"
	"admin/service"
	"admin/service/dto"
	"github.com/kingwel-xie/k2/common/api"
)

type SysConfig struct {
	api.Api
}

// GetPage 获取配置管理列表
// @Summary 获取配置管理列表
// @Description 获取配置管理列表
// @Tags 配置管理
// @Param configName query string false "名称"
// @Param configKey query string false "key"
// @Param configType query string false "类型"
// @Param isFrontend query int false "是否前端"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/config [get]
// @Security Bearer
func (e SysConfig) GetPage(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}

	list := make([]models.SysConfig, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取配置管理
// @Summary 获取配置管理
// @Description 获取配置管理
// @Tags 配置管理
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SysConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/config/{id} [get]
// @Security Bearer
func (e SysConfig) Get(c *gin.Context) {
	req := dto.SysConfigGetReq{}
	s := service.SysConfig{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}
	var object models.SysConfig

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建配置管理
// @Summary 创建配置管理
// @Description 创建配置管理
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysConfigControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "创建成功"}"
// @Router /api/v1/config [post]
// @Security Bearer
func (e SysConfig) Insert(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigControl{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改配置管理
// @Summary 修改配置管理
// @Description 修改配置管理
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param id path string true "id"
// @Param data body dto.SysConfigControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/config/{id} [put]
// @Security Bearer
func (e SysConfig) Update(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigControl{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除配置管理
// @Summary 删除配置管理
// @Description 删除配置管理
// @Tags 配置管理
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/config [delete]
// @Security Bearer
func (e SysConfig) Delete(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Get2SysApp 获取系统配置信息
// @Summary 获取系统前台配置信息。此接口不验证权限
// @Description 获取系统配置信息。此接口不验证权限
// @Tags 配置管理
// @Success 200 {object} response.Response{data=map[string]string} "{"code": 200, "data": [...]}"
// @Router /api/v1/app-config [get]
func (e SysConfig) Get2SysApp(c *gin.Context) {
	req := dto.SysConfigGetToSysAppReq{}
	s := service.SysConfig{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}
	// 控制只读前台的数据
	req.IsFrontend = 1
	list := make([]models.SysConfig, 0)
	err = s.GetWithKeyList(&req, &list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	mp := make(map[string]string)
	for i := 0; i < len(list); i++ {
		key := list[i].ConfigKey
		if key != "" {
			mp[key] = list[i].ConfigValue
		}
	}
	e.OK(mp, "查询成功")
}

// Get2Set 获取配置
// @Summary 获取配置
// @Description 界面操作设置配置值的获取
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Success 200 {object} response.Response{data=map[string]interface{}}	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/set-config [get]
// @Security Bearer
func (e SysConfig) Get2Set(c *gin.Context) {
	s := service.SysConfig{}
	req := make([]dto.GetSetSysConfigReq, 0)
	err := e.MakeContext(c).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}
	err = s.GetForSet(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	m := make(map[string]interface{}, 0)
	for _, v := range req {
		m[v.ConfigKey] = v.ConfigValue
	}
	e.OK(m, "查询成功")
}

// Update2Set 设置配置
// @Summary 设置配置
// @Description 界面操作设置配置值
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body []dto.GetSetSysConfigReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/set-config [put]
// @Security Bearer
func (e SysConfig) Update2Set(c *gin.Context) {
	s := service.SysConfig{}
	req := make([]dto.GetSetSysConfigReq, 0)
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}

	err = s.UpdateForSet(&req)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}

	e.OK("", "更新成功")
}

// GetSysConfigByKEYForService 根据Key获取SysConfig的Service
// @Summary 根据Key获取SysConfig的Service
// @Description 根据Key获取SysConfig的Service
// @Tags 配置管理
// @Param configKey path string false "configKey"
// @Success 200 {object} response.Response{data=models.SysConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/configKey/{configKey} [get]
// @Security Bearer
func (e SysConfig) GetSysConfigByKEYForService(c *gin.Context) {
	var s = new(service.SysConfig)
	var req = new(dto.SysConfigByKeyReq)
	var resp = new(dto.GetSysConfigByKEYForServiceResp)
	err := e.MakeContext(c).
		Bind(req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(400, err, "参数错误")
		return
	}

	err = s.GetWithKey(req, resp)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(resp, "Ok")
}
