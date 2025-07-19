package handler

import (
	"strconv"
	"time"

	"github.com/kingwrcy/moments/db"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type ChildHandler struct {
	injector do.Injector
}

func NewChildHandler(injector do.Injector) *ChildHandler {
	return &ChildHandler{injector: injector}
}

// 获取儿童档案列表
func (h *ChildHandler) List(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	userId := c.Get("userId").(int32)
	var children []db.Child

	if err := tx.Where("userId = ?", userId).Find(&children).Error; err != nil {
		log.Error().Msgf("获取儿童档案列表失败: %s", err)
		return FailRespWithMsg(c, Fail, "获取儿童档案列表失败")
	}

	return SuccessResp(c, children)
}

// 获取儿童档案详情
func (h *ChildHandler) Get(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var child db.Child
	if err := tx.Where("id = ?", id).First(&child).Error; err != nil {
		log.Error().Msgf("获取儿童档案详情失败: %s", err)
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}

	return SuccessResp(c, child)
}

// 创建儿童档案
func (h *ChildHandler) Create(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	var child db.Child
	if err := c.Bind(&child); err != nil {
		return FailResp(c, ParamError)
	}

	userId := c.Get("userId").(int32)
	child.UserId = userId
	now := time.Now()
	child.CreatedAt = &now
	child.UpdatedAt = &now

	if err := tx.Create(&child).Error; err != nil {
		log.Error().Msgf("创建儿童档案失败: %s", err)
		return FailRespWithMsg(c, Fail, "创建儿童档案失败")
	}

	return SuccessResp(c, child)
}

// 更新儿童档案
func (h *ChildHandler) Update(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var child db.Child
	if err := tx.Where("id = ?", id).First(&child).Error; err != nil {
		log.Error().Msgf("获取儿童档案失败: %s", err)
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}

	var updateData db.Child
	if err := c.Bind(&updateData); err != nil {
		return FailResp(c, ParamError)
	}

	now := time.Now()
	child.UpdatedAt = &now
	child.Name = updateData.Name
	child.Gender = updateData.Gender
	child.BirthDate = updateData.BirthDate
	child.BloodType = updateData.BloodType
	child.Avatar = updateData.Avatar
	child.Description = updateData.Description

	if err := tx.Save(&child).Error; err != nil {
		log.Error().Msgf("更新儿童档案失败: %s", err)
		return FailRespWithMsg(c, Fail, "更新儿童档案失败")
	}

	return SuccessResp(c, child)
}

// 删除儿童档案
func (h *ChildHandler) Delete(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	if err := tx.Where("id = ?", id).Delete(&db.Child{}).Error; err != nil {
		log.Error().Msgf("删除儿童档案失败: %s", err)
		return FailRespWithMsg(c, Fail, "删除儿童档案失败")
	}

	return SuccessResp(c, map[string]any{"message": "删除成功"})
}