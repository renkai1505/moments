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

type GrowthHandler struct {
	injector do.Injector
}

func NewGrowthHandler(injector do.Injector) *GrowthHandler {
	return &GrowthHandler{injector: injector}
}

// 获取成长记录列表
func (h *GrowthHandler) List(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	childId, err := strconv.Atoi(c.QueryParam("childId"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var records []db.GrowthRecord
	if err := tx.Where("childId = ?", childId).Order("recordDate DESC").Find(&records).Error; err != nil {
		log.Error().Msgf("获取成长记录列表失败: %s", err)
		return FailRespWithMsg(c, Fail, "获取成长记录列表失败")
	}

	return SuccessResp(c, records)
}

// 获取成长记录详情
func (h *GrowthHandler) Get(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var record db.GrowthRecord
	if err := tx.Where("id = ?", id).First(&record).Error; err != nil {
		log.Error().Msgf("获取成长记录详情失败: %s", err)
		return FailRespWithMsg(c, Fail, "成长记录不存在")
	}

	return SuccessResp(c, record)
}

// 创建成长记录
func (h *GrowthHandler) Create(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	var record db.GrowthRecord
	if err := c.Bind(&record); err != nil {
		return FailResp(c, ParamError)
	}

	now := time.Now()
	record.CreatedAt = &now
	record.UpdatedAt = &now

	if err := tx.Create(&record).Error; err != nil {
		log.Error().Msgf("创建成长记录失败: %s", err)
		return FailRespWithMsg(c, Fail, "创建成长记录失败")
	}

	return SuccessResp(c, record)
}

// 更新成长记录
func (h *GrowthHandler) Update(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	var record db.GrowthRecord
	if err := tx.Where("id = ?", id).First(&record).Error; err != nil {
		log.Error().Msgf("获取成长记录失败: %s", err)
		return FailRespWithMsg(c, Fail, "成长记录不存在")
	}

	var updateData db.GrowthRecord
	if err := c.Bind(&updateData); err != nil {
		return FailResp(c, ParamError)
	}

	now := time.Now()
	record.UpdatedAt = &now
	record.Height = updateData.Height
	record.Weight = updateData.Weight
	record.HeadCirc = updateData.HeadCirc
	record.RecordDate = updateData.RecordDate
	record.Notes = updateData.Notes

	if err := tx.Save(&record).Error; err != nil {
		log.Error().Msgf("更新成长记录失败: %s", err)
		return FailRespWithMsg(c, Fail, "更新成长记录失败")
	}

	return SuccessResp(c, record)
}

// 删除成长记录
func (h *GrowthHandler) Delete(c echo.Context) error {
	log := do.MustInvoke[zerolog.Logger](h.injector)
	tx := do.MustInvoke[*gorm.DB](h.injector)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailResp(c, ParamError)
	}

	if err := tx.Where("id = ?", id).Delete(&db.GrowthRecord{}).Error; err != nil {
		log.Error().Msgf("删除成长记录失败: %s", err)
		return FailRespWithMsg(c, Fail, "删除成长记录失败")
	}

	return SuccessResp(c, map[string]any{"message": "删除成功"})
}