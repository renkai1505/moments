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

type MilestoneHandler struct {
	injector do.Injector
}

func NewMilestoneHandler(injector do.Injector) *MilestoneHandler {
	return &MilestoneHandler{injector: injector}
}

// 获取里程碑列表
func (h *MilestoneHandler) List(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	childId, err := strconv.Atoi(c.QueryParam("childId"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var milestones []db.Milestone
	if err := tx.Where("childId = ?", childId).Order("date DESC").Find(&milestones).Error; err != nil {
		log.Error().Msgf("获取里程碑列表失败: %s", err)
		return FailRespWithMsg(c, Fail, "获取里程碑列表失败")
	}

	return SuccessResp(c, milestones)
}

// 获取里程碑详情
func (h *MilestoneHandler) Get(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var milestone db.Milestone
	if err := tx.Where("id = ?", id).First(&milestone).Error; err != nil {
		log.Error().Msgf("获取里程碑详情失败: %s", err)
		return FailRespWithMsg(c, Fail, "里程碑不存在")
	}

	return SuccessResp(c, milestone)
}

// 创建里程碑
func (h *MilestoneHandler) Create(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	var milestone db.Milestone
	if err := c.Bind(&milestone); err != nil {
		return FailResp(c, ParamError)
	}

	now := time.Now()
	milestone.CreatedAt = &now
	milestone.UpdatedAt = &now

	if err := tx.Create(&milestone).Error; err != nil {
		log.Error().Msgf("创建里程碑失败: %s", err)
		return FailRespWithMsg(c, Fail, "创建里程碑失败")
	}

	return SuccessResp(c, milestone)
}

// 更新里程碑
func (h *MilestoneHandler) Update(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var milestone db.Milestone
	if err := tx.Where("id = ?", id).First(&milestone).Error; err != nil {
		log.Error().Msgf("获取里程碑失败: %s", err)
		return FailRespWithMsg(c, Fail, "里程碑不存在")
	}

	var updateData db.Milestone
	if err := c.Bind(&updateData); err != nil {
		return FailResp(c, ParamError)
	}

	now := time.Now()
	milestone.UpdatedAt = &now
	milestone.Title = updateData.Title
	milestone.Description = updateData.Description
	milestone.Category = updateData.Category
	milestone.Date = updateData.Date
	milestone.IsImportant = updateData.IsImportant

	if err := tx.Save(&milestone).Error; err != nil {
		log.Error().Msgf("更新里程碑失败: %s", err)
		return FailRespWithMsg(c, Fail, "更新里程碑失败")
	}

	return SuccessResp(c, milestone)
}

// 删除里程碑
func (h *MilestoneHandler) Delete(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	if err := tx.Where("id = ?", id).Delete(&db.Milestone{}).Error; err != nil {
		log.Error().Msgf("删除里程碑失败: %s", err)
		return FailRespWithMsg(c, Fail, "删除里程碑失败")
	}

	return SuccessResp(c, map[string]any{"message": "删除成功"})
}