package vo

import "time"

// ChildVO 儿童视图对象
type ChildVO struct {
	Id          int32      `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Nickname    string     `json:"nickname,omitempty"`
	BirthDate   *time.Time `json:"birthDate,omitempty"`
	Age         int        `json:"age,omitempty"`         // 计算出的年龄
	AgeInDays   int        `json:"ageInDays,omitempty"`   // 出生天数
	Gender      string     `json:"gender,omitempty"`
	AvatarUrl   string     `json:"avatarUrl,omitempty"`
	CoverUrl    string     `json:"coverUrl,omitempty"`
	Height      float32    `json:"height,omitempty"`
	Weight      float32    `json:"weight,omitempty"`
	BloodType   string     `json:"bloodType,omitempty"`
	Hobbies     string     `json:"hobbies,omitempty"`
	Description string     `json:"description,omitempty"`
	ParentId    int32      `json:"parentId,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	Parent      *UserVO    `json:"parent,omitempty"`
}

// SaveChildReq 保存儿童档案请求
type SaveChildReq struct {
	Id          int32      `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Nickname    string     `json:"nickname,omitempty"`
	BirthDate   *time.Time `json:"birthDate,omitempty"`
	Gender      string     `json:"gender,omitempty"`
	AvatarUrl   string     `json:"avatarUrl,omitempty"`
	CoverUrl    string     `json:"coverUrl,omitempty"`
	Height      float32    `json:"height,omitempty"`
	Weight      float32    `json:"weight,omitempty"`
	BloodType   string     `json:"bloodType,omitempty"`
	Hobbies     string     `json:"hobbies,omitempty"`
	Description string     `json:"description,omitempty"`
}

// ListChildReq 查询儿童列表请求
type ListChildReq struct {
	Page     int `json:"page,omitempty"`
	Size     int `json:"size,omitempty"`
	ParentId int `json:"parentId,omitempty"`
}

// GrowthRecordVO 成长记录视图对象
type GrowthRecordVO struct {
	Id          int32              `json:"id,omitempty"`
	ChildId     int32              `json:"childId,omitempty"`
	Title       string             `json:"title,omitempty"`
	Content     string             `json:"content,omitempty"`
	RecordType  string             `json:"recordType,omitempty"`
	Height      *float32           `json:"height,omitempty"`
	Weight      *float32           `json:"weight,omitempty"`
	Imgs        string             `json:"imgs,omitempty"`
	Location    string             `json:"location,omitempty"`
	RecordDate  *time.Time         `json:"recordDate,omitempty"`
	Milestone   string             `json:"milestone,omitempty"`
	Mood        string             `json:"mood,omitempty"`
	Weather     string             `json:"weather,omitempty"`
	Tags        string             `json:"tags,omitempty"`
	ParentId    int32              `json:"parentId,omitempty"`
	Pinned      *bool              `json:"pinned,omitempty"`
	ShowType    *int32             `json:"showType,omitempty"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty"`
	Child       *ChildVO           `json:"child,omitempty"`
	Parent      *UserVO            `json:"parent,omitempty"`
	ImgConfigs  *[]map[string]any  `json:"imgConfigs,omitempty"`
}

// SaveGrowthRecordReq 保存成长记录请求
type SaveGrowthRecordReq struct {
	Id          int32      `json:"id,omitempty"`
	ChildId     int32      `json:"childId,omitempty"`
	Title       string     `json:"title,omitempty"`
	Content     string     `json:"content,omitempty"`
	RecordType  string     `json:"recordType,omitempty"`
	Height      *float32   `json:"height,omitempty"`
	Weight      *float32   `json:"weight,omitempty"`
	Imgs        []string   `json:"imgs,omitempty"`
	Location    string     `json:"location,omitempty"`
	RecordDate  *time.Time `json:"recordDate,omitempty"`
	Milestone   string     `json:"milestone,omitempty"`
	Mood        string     `json:"mood,omitempty"`
	Weather     string     `json:"weather,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	Pinned      *bool      `json:"pinned,omitempty"`
	ShowType    *int32     `json:"showType,omitempty"`
}

// ListGrowthRecordReq 查询成长记录请求
type ListGrowthRecordReq struct {
	Page       int        `json:"page,omitempty"`
	Size       int        `json:"size,omitempty"`
	ChildId    int32      `json:"childId,omitempty"`
	RecordType string     `json:"recordType,omitempty"`
	Start      *time.Time `json:"start,omitempty"`
	End        *time.Time `json:"end,omitempty"`
	ParentId   int32      `json:"parentId,omitempty"`
}

// TimelineItemVO 时间轴项目
type TimelineItemVO struct {
	Date    string               `json:"date,omitempty"`    // 日期格式：YYYY-MM-DD
	Records []GrowthRecordVO     `json:"records,omitempty"` // 该日期的记录
}

// GrowthStatsVO 成长统计
type GrowthStatsVO struct {
	ChildId      int32                 `json:"childId,omitempty"`
	TotalRecords int                   `json:"totalRecords,omitempty"`
	RecordTypes  map[string]int        `json:"recordTypes,omitempty"`  // 各类型记录数量
	HeightGrowth []GrowthDataPoint     `json:"heightGrowth,omitempty"` // 身高增长曲线
	WeightGrowth []GrowthDataPoint     `json:"weightGrowth,omitempty"` // 体重增长曲线
	Milestones   []MilestoneVO         `json:"milestones,omitempty"`   // 里程碑事件
}

// GrowthDataPoint 增长数据点
type GrowthDataPoint struct {
	Date  string  `json:"date,omitempty"`
	Value float32 `json:"value,omitempty"`
}

// MilestoneVO 里程碑事件
type MilestoneVO struct {
	Id          int32      `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	AgeAtTime   string     `json:"ageAtTime,omitempty"` // 当时的年龄描述
}