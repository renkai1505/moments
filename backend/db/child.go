package db

import (
	"time"
)

// Child 儿童档案
type Child struct {
	Id          int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	Name        string     `gorm:"column:name;NOT NULL" json:"name,omitempty"`
	Gender      string     `gorm:"column:gender" json:"gender,omitempty"` // 性别：男/女
	BirthDate   *time.Time `gorm:"column:birthDate" json:"birthDate,omitempty"`
	BloodType   string     `gorm:"column:bloodType" json:"bloodType,omitempty"` // 血型
	Avatar      string     `gorm:"column:avatar" json:"avatar,omitempty"`       // 头像
	Description string     `gorm:"column:description" json:"description,omitempty"` // 描述
	UserId      int32      `gorm:"column:userId;NOT NULL" json:"userId,omitempty"`
	CreatedAt   *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
	User        *User      `json:"user,omitempty"`
}

func (c *Child) TableName() string {
	return "Child"
}

// GrowthRecord 成长记录
type GrowthRecord struct {
	Id        int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	ChildId   int32      `gorm:"column:childId;NOT NULL" json:"childId,omitempty"`
	Height    float64    `gorm:"column:height" json:"height,omitempty"`       // 身高(cm)
	Weight    float64    `gorm:"column:weight" json:"weight,omitempty"`       // 体重(kg)
	HeadCirc  float64    `gorm:"column:headCirc" json:"headCirc,omitempty"`   // 头围(cm)
	RecordDate *time.Time `gorm:"column:recordDate;NOT NULL" json:"recordDate,omitempty"`
	Notes     string     `gorm:"column:notes" json:"notes,omitempty"`         // 备注
	CreatedAt *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
	Child     *Child     `json:"child,omitempty"`
}

func (g *GrowthRecord) TableName() string {
	return "GrowthRecord"
}

// Milestone 成长里程碑
type Milestone struct {
	Id          int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	ChildId     int32      `gorm:"column:childId;NOT NULL" json:"childId,omitempty"`
	Title       string     `gorm:"column:title;NOT NULL" json:"title,omitempty"`
	Description string     `gorm:"column:description" json:"description,omitempty"`
	Category    string     `gorm:"column:category" json:"category,omitempty"` // 分类：运动/语言/认知/社交等
	Date        *time.Time `gorm:"column:date;NOT NULL" json:"date,omitempty"`
	IsImportant bool       `gorm:"column:isImportant;default:false" json:"isImportant,omitempty"` // 是否重要里程碑
	CreatedAt   *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`
	Child       *Child     `json:"child,omitempty"`
}

func (m *Milestone) TableName() string {
	return "Milestone"
}