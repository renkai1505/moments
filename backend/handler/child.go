package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/util"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

type ChildHandler struct {
	BaseHandler
}

func NewChildHandler(injector do.Injector) (*ChildHandler, error) {
	base, err := NewBaseHandler(injector)
	if err != nil {
		return nil, err
	}
	return &ChildHandler{BaseHandler: base}, nil
}

// SaveChild 保存儿童档案
func (h *ChildHandler) SaveChild(c echo.Context) error {
	var req vo.SaveChildReq
	if err := c.Bind(&req); err != nil {
		h.log.Error().Err(err).Msg("绑定参数失败")
		return FailRespWithMsg(c, ParamError, "参数错误")
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	var child db.Child
	if req.Id > 0 {
		// 更新
		if err := h.db.First(&child, req.Id).Error; err != nil {
			return FailRespWithMsg(c, Fail, "儿童档案不存在")
		}
		if child.ParentId != currentUser.Id && currentUser.Id != 1 {
			return FailRespWithMsg(c, Fail, "无权限操作")
		}
	} else {
		// 新增
		child.ParentId = currentUser.Id
		now := time.Now()
		child.CreatedAt = &now
	}

	child.Name = req.Name
	child.Nickname = req.Nickname
	child.BirthDate = req.BirthDate
	child.Gender = req.Gender
	child.AvatarUrl = req.AvatarUrl
	child.CoverUrl = req.CoverUrl
	child.Height = req.Height
	child.Weight = req.Weight
	child.BloodType = req.BloodType
	child.Hobbies = req.Hobbies
	child.Description = req.Description

	now := time.Now()
	child.UpdatedAt = &now

	if err := h.db.Save(&child).Error; err != nil {
		h.log.Error().Err(err).Msg("保存儿童档案失败")
		return FailRespWithMsg(c, Fail, "保存失败")
	}

	return SuccessResp(c, child.Id)
}

// ListChildren 获取儿童列表
func (h *ChildHandler) ListChildren(c echo.Context) error {
	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	var children []db.Child
	query := h.db.Where("parentId = ?", currentUser.Id)

	if err := query.Preload("Parent").Find(&children).Error; err != nil {
		h.log.Error().Err(err).Msg("查询儿童列表失败")
		return FailRespWithMsg(c, Fail, "查询失败")
	}

	// 转换为VO并计算年龄
	var childVOs []vo.ChildVO
	for _, child := range children {
		childVO := h.convertChildToVO(child)
		childVOs = append(childVOs, childVO)
	}

	return SuccessResp(c, childVOs)
}

// GetChild 获取儿童详情
func (h *ChildHandler) GetChild(c echo.Context) error {
	childId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailRespWithMsg(c, ParamError, "参数错误")
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	var child db.Child
	if err := h.db.Preload("Parent").First(&child, childId).Error; err != nil {
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}

	if child.ParentId != currentUser.Id && currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	childVO := h.convertChildToVO(child)
	return SuccessResp(c, childVO)
}

// DeleteChild 删除儿童档案
func (h *ChildHandler) DeleteChild(c echo.Context) error {
	childId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return FailRespWithMsg(c, ParamError, "参数错误")
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	var child db.Child
	if err := h.db.First(&child, childId).Error; err != nil {
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}

	if child.ParentId != currentUser.Id && currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "无权限操作")
	}

	// 删除相关的成长记录
	if err := h.db.Where("childId = ?", childId).Delete(&db.GrowthRecord{}).Error; err != nil {
		h.log.Error().Err(err).Msg("删除成长记录失败")
	}

	// 删除儿童档案
	if err := h.db.Delete(&child).Error; err != nil {
		h.log.Error().Err(err).Msg("删除儿童档案失败")
		return FailRespWithMsg(c, Fail, "删除失败")
	}

	return SuccessResp(c, map[string]interface{}{})
}

// SaveGrowthRecord 保存成长记录
func (h *ChildHandler) SaveGrowthRecord(c echo.Context) error {
	var req vo.SaveGrowthRecordReq
	if err := c.Bind(&req); err != nil {
		h.log.Error().Err(err).Msg("绑定参数失败")
		return FailRespWithMsg(c, ParamError, "参数错误")
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	// 验证儿童存在且有权限
	var child db.Child
	if err := h.db.First(&child, req.ChildId).Error; err != nil {
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}
	if child.ParentId != currentUser.Id && currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "无权限操作")
	}

	var record db.GrowthRecord
	if req.Id > 0 {
		// 更新
		if err := h.db.First(&record, req.Id).Error; err != nil {
			return FailRespWithMsg(c, Fail, "成长记录不存在")
		}
		if record.ParentId != currentUser.Id && currentUser.Id != 1 {
			return FailRespWithMsg(c, Fail, "无权限操作")
		}
	} else {
		// 新增
		record.ParentId = currentUser.Id
		now := time.Now()
		record.CreatedAt = &now
	}

	record.ChildId = req.ChildId
	record.Title = req.Title
	record.Content = req.Content
	record.RecordType = req.RecordType
	record.Height = req.Height
	record.Weight = req.Weight
	record.Location = req.Location
	record.RecordDate = req.RecordDate
	record.Milestone = req.Milestone
	record.Mood = req.Mood
	record.Weather = req.Weather
	record.Pinned = req.Pinned
	record.ShowType = req.ShowType

	// 处理图片
	if len(req.Imgs) > 0 {
		imgBytes, err := json.Marshal(req.Imgs)
		if err == nil {
			record.Imgs = string(imgBytes)
		}
	}

	// 处理标签
	if len(req.Tags) > 0 {
		record.Tags = strings.Join(req.Tags, ",")
	}

	now := time.Now()
	record.UpdatedAt = &now

	if err := h.db.Save(&record).Error; err != nil {
		h.log.Error().Err(err).Msg("保存成长记录失败")
		return FailRespWithMsg(c, Fail, "保存失败")
	}

	return SuccessResp(c, record.Id)
}

// ListGrowthRecords 获取成长记录列表
func (h *ChildHandler) ListGrowthRecords(c echo.Context) error {
	var req vo.ListGrowthRecordReq
	if err := c.Bind(&req); err != nil {
		// 如果是GET请求，尝试从查询参数获取
		if c.Request().Method == "GET" {
			req.Page, _ = strconv.Atoi(c.QueryParam("page"))
			if req.Page <= 0 {
				req.Page = 1
			}
			req.Size, _ = strconv.Atoi(c.QueryParam("size"))
			if req.Size <= 0 {
				req.Size = 10
			}
			childId, _ := strconv.Atoi(c.QueryParam("childId"))
			req.ChildId = int32(childId)
			req.RecordType = c.QueryParam("recordType")
		} else {
			h.log.Error().Err(err).Msg("绑定参数失败")
			return FailRespWithMsg(c, ParamError, "参数错误")
		}
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	query := h.db.Model(&db.GrowthRecord{})

	// 权限控制：只能查看自己的儿童记录
	if currentUser.Id != 1 {
		query = query.Where("parentId = ?", currentUser.Id)
	}

	if req.ChildId > 0 {
		query = query.Where("childId = ?", req.ChildId)
	}
	if req.RecordType != "" {
		query = query.Where("recordType = ?", req.RecordType)
	}
	if req.Start != nil {
		query = query.Where("recordDate >= ?", req.Start)
	}
	if req.End != nil {
		query = query.Where("recordDate <= ?", req.End)
	}

	var total int64
	query.Count(&total)

	var records []db.GrowthRecord
	offset := (req.Page - 1) * req.Size
	if err := query.Preload("Child").Preload("Parent").
		Order("recordDate DESC, createdAt DESC").
		Offset(offset).Limit(req.Size).Find(&records).Error; err != nil {
		h.log.Error().Err(err).Msg("查询成长记录失败")
		return FailRespWithMsg(c, Fail, "查询失败")
	}

	// 转换为VO
	var recordVOs []vo.GrowthRecordVO
	for _, record := range records {
		recordVO := h.convertGrowthRecordToVO(record)
		recordVOs = append(recordVOs, recordVO)
	}

	hasNext := int64((req.Page)*req.Size) < total

	result := map[string]interface{}{
		"list":    recordVOs,
		"total":   total,
		"hasNext": hasNext,
	}

	return SuccessResp(c, result)
}

// GetGrowthTimeline 获取成长时间轴
func (h *ChildHandler) GetGrowthTimeline(c echo.Context) error {
	childId, err := strconv.Atoi(c.Param("childId"))
	if err != nil {
		return FailRespWithMsg(c, ParamError, "参数错误")
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	// 验证权限
	var child db.Child
	if err := h.db.First(&child, childId).Error; err != nil {
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}
	if child.ParentId != currentUser.Id && currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	var records []db.GrowthRecord
	if err := h.db.Where("childId = ?", childId).
		Preload("Child").Preload("Parent").
		Order("recordDate DESC, createdAt DESC").
		Find(&records).Error; err != nil {
		h.log.Error().Err(err).Msg("查询成长记录失败")
		return FailRespWithMsg(c, Fail, "查询失败")
	}

	// 按日期分组
	timelineMap := make(map[string][]vo.GrowthRecordVO)
	for _, record := range records {
		recordVO := h.convertGrowthRecordToVO(record)
		date := ""
		if record.RecordDate != nil {
			date = record.RecordDate.Format("2006-01-02")
		} else if record.CreatedAt != nil {
			date = record.CreatedAt.Format("2006-01-02")
		}
		timelineMap[date] = append(timelineMap[date], recordVO)
	}

	// 转换为时间轴格式
	var timeline []vo.TimelineItemVO
	for date, records := range timelineMap {
		timeline = append(timeline, vo.TimelineItemVO{
			Date:    date,
			Records: records,
		})
	}

	return SuccessResp(c, timeline)
}

// GetGrowthStats 获取成长统计
func (h *ChildHandler) GetGrowthStats(c echo.Context) error {
	childId, err := strconv.Atoi(c.Param("childId"))
	if err != nil {
		return FailRespWithMsg(c, ParamError, "参数错误")
	}

	ctx := c.(*CustomContext)
	currentUser := ctx.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, TokenInvalid, "用户未登录")
	}

	// 验证权限
	var child db.Child
	if err := h.db.First(&child, childId).Error; err != nil {
		return FailRespWithMsg(c, Fail, "儿童档案不存在")
	}
	if child.ParentId != currentUser.Id && currentUser.Id != 1 {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	stats := vo.GrowthStatsVO{
		ChildId:     int32(childId),
		RecordTypes: make(map[string]int),
	}

	// 统计总记录数
	var totalRecords int64
	h.db.Model(&db.GrowthRecord{}).Where("childId = ?", childId).Count(&totalRecords)
	stats.TotalRecords = int(totalRecords)

	// 统计各类型记录数量
	var typeStats []struct {
		RecordType string
		Count      int
	}
	h.db.Model(&db.GrowthRecord{}).
		Select("recordType, count(*) as count").
		Where("childId = ?", childId).
		Group("recordType").
		Find(&typeStats)

	for _, stat := range typeStats {
		stats.RecordTypes[stat.RecordType] = stat.Count
	}

	// 获取身高体重增长数据
	var growthRecords []db.GrowthRecord
	h.db.Where("childId = ? AND (height IS NOT NULL OR weight IS NOT NULL)", childId).
		Order("recordDate ASC").
		Find(&growthRecords)

	for _, record := range growthRecords {
		date := ""
		if record.RecordDate != nil {
			date = record.RecordDate.Format("2006-01-02")
		}
		if record.Height != nil && *record.Height > 0 {
			stats.HeightGrowth = append(stats.HeightGrowth, vo.GrowthDataPoint{
				Date:  date,
				Value: *record.Height,
			})
		}
		if record.Weight != nil && *record.Weight > 0 {
			stats.WeightGrowth = append(stats.WeightGrowth, vo.GrowthDataPoint{
				Date:  date,
				Value: *record.Weight,
			})
		}
	}

	// 获取里程碑事件
	var milestoneRecords []db.GrowthRecord
	h.db.Where("childId = ? AND recordType = ? AND milestone != ''", childId, "milestone").
		Order("recordDate ASC").
		Find(&milestoneRecords)

	for _, record := range milestoneRecords {
		ageAtTime := ""
		if child.BirthDate != nil && record.RecordDate != nil {
			days := int(record.RecordDate.Sub(*child.BirthDate).Hours() / 24)
			if days > 365 {
				years := days / 365
				remainingDays := days % 365
				ageAtTime = fmt.Sprintf("%d岁%d天", years, remainingDays)
			} else {
				ageAtTime = fmt.Sprintf("%d天", days)
			}
		}

		stats.Milestones = append(stats.Milestones, vo.MilestoneVO{
			Id:          record.Id,
			Title:       record.Title,
			Description: record.Milestone,
			Date:        record.RecordDate,
			AgeAtTime:   ageAtTime,
		})
	}

	return SuccessResp(c, stats)
}

// convertChildToVO 转换儿童为VO
func (h *ChildHandler) convertChildToVO(child db.Child) vo.ChildVO {
	childVO := vo.ChildVO{
		Id:          child.Id,
		Name:        child.Name,
		Nickname:    child.Nickname,
		BirthDate:   child.BirthDate,
		Gender:      child.Gender,
		AvatarUrl:   child.AvatarUrl,
		CoverUrl:    child.CoverUrl,
		Height:      child.Height,
		Weight:      child.Weight,
		BloodType:   child.BloodType,
		Hobbies:     child.Hobbies,
		Description: child.Description,
		ParentId:    child.ParentId,
		CreatedAt:   child.CreatedAt,
		UpdatedAt:   child.UpdatedAt,
	}

	// 计算年龄
	if child.BirthDate != nil {
		now := time.Now()
		childVO.AgeInDays = int(now.Sub(*child.BirthDate).Hours() / 24)
		childVO.Age = now.Year() - child.BirthDate.Year()
		if now.YearDay() < child.BirthDate.YearDay() {
			childVO.Age--
		}
	}

	if child.Parent != nil {
		childVO.Parent = &vo.UserVO{
			Id:        child.Parent.Id,
			Username:  child.Parent.Username,
			Nickname:  child.Parent.Nickname,
			AvatarUrl: child.Parent.AvatarUrl,
			Slogan:    child.Parent.Slogan,
			CoverUrl:  child.Parent.CoverUrl,
			Email:     child.Parent.Email,
		}
	}

	return childVO
}

// convertGrowthRecordToVO 转换成长记录为VO
func (h *ChildHandler) convertGrowthRecordToVO(record db.GrowthRecord) vo.GrowthRecordVO {
	recordVO := vo.GrowthRecordVO{
		Id:          record.Id,
		ChildId:     record.ChildId,
		Title:       record.Title,
		Content:     record.Content,
		RecordType:  record.RecordType,
		Height:      record.Height,
		Weight:      record.Weight,
		Imgs:        record.Imgs,
		Location:    record.Location,
		RecordDate:  record.RecordDate,
		Milestone:   record.Milestone,
		Mood:        record.Mood,
		Weather:     record.Weather,
		Tags:        record.Tags,
		ParentId:    record.ParentId,
		Pinned:      record.Pinned,
		ShowType:    record.ShowType,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}

	// 处理图片配置
	if record.Imgs != "" {
		var imgs []string
		if err := json.Unmarshal([]byte(record.Imgs), &imgs); err == nil {
			imgConfigs := make([]map[string]any, len(imgs))
			for i, img := range imgs {
				imgConfigs[i] = map[string]any{
					"url":      img,
					"thumbUrl": util.GetThumbUrl(img),
				}
			}
			recordVO.ImgConfigs = &imgConfigs
		}
	}

	// 关联数据
	if record.Child != nil {
		recordVO.Child = &vo.ChildVO{
			Id:        record.Child.Id,
			Name:      record.Child.Name,
			Nickname:  record.Child.Nickname,
			BirthDate: record.Child.BirthDate,
			Gender:    record.Child.Gender,
			AvatarUrl: record.Child.AvatarUrl,
		}
	}

	if record.Parent != nil {
		recordVO.Parent = &vo.UserVO{
			Id:        record.Parent.Id,
			Username:  record.Parent.Username,
			Nickname:  record.Parent.Nickname,
			AvatarUrl: record.Parent.AvatarUrl,
			Slogan:    record.Parent.Slogan,
			CoverUrl:  record.Parent.CoverUrl,
			Email:     record.Parent.Email,
		}
	}

	return recordVO
}