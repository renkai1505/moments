package db

import (
	"time"
)

// Child 儿童档案模型
type Child struct {
	Id          int32      `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`                             //儿童ID
	Name        string     `gorm:"column:name;NOT NULL" json:"name,omitempty"`                                     //姓名
	Nickname    string     `gorm:"column:nickname" json:"nickname,omitempty"`                                      //昵称
	BirthDate   *time.Time `gorm:"column:birthDate" json:"birthDate,omitempty"`                                    //出生日期
	Gender      string     `gorm:"column:gender" json:"gender,omitempty"`                                          //性别 M/F
	AvatarUrl   string     `gorm:"column:avatarUrl" json:"avatarUrl,omitempty"`                                    //头像URL
	CoverUrl    string     `gorm:"column:coverUrl" json:"coverUrl,omitempty"`                                      //封面URL
	Height      float32    `gorm:"column:height" json:"height,omitempty"`                                          //身高(cm)
	Weight      float32    `gorm:"column:weight" json:"weight,omitempty"`                                          //体重(kg)
	BloodType   string     `gorm:"column:bloodType" json:"bloodType,omitempty"`                                    //血型
	Hobbies     string     `gorm:"column:hobbies" json:"hobbies,omitempty"`                                        //爱好
	Description string     `gorm:"column:description" json:"description,omitempty"`                               //描述
	ParentId    int32      `gorm:"column:parentId;NOT NULL" json:"parentId,omitempty"`                             //家长ID
	CreatedAt   *time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"` //创建时间
	UpdatedAt   *time.Time `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`                           //更新时间

	Parent      *User           `gorm:"foreignKey:ParentId" json:"parent,omitempty"`                           //关联的家长
	GrowthRecords []GrowthRecord `gorm:"foreignKey:ChildId" json:"growthRecords,omitempty"`               //关联的成长记录

}

func (c *Child) TableName() string {
	return "Child"
}

// GrowthRecord 成长记录模型（基于现有Memo扩展）
type GrowthRecord struct {
	Id          int32              `gorm:"column:id;primary_key;NOT NULL" json:"id,omitempty"`
	ChildId     int32              `gorm:"column:childId;NOT NULL" json:"childId,omitempty"`                           //儿童ID
	Title       string             `gorm:"column:title" json:"title,omitempty"`                                        //标题
	Content     string             `gorm:"column:content" json:"content,omitempty"`                                    //内容
	RecordType  string             `gorm:"column:recordType" json:"recordType,omitempty"`                              //记录类型: growth/health/study/play/milestone
	Height      *float32           `gorm:"column:height" json:"height,omitempty"`                                      //身高记录
	Weight      *float32           `gorm:"column:weight" json:"weight,omitempty"`                                      //体重记录
	Imgs        string             `gorm:"column:imgs" json:"imgs,omitempty"`                                          //图片
	Location    string             `gorm:"column:location" json:"location,omitempty"`                                  //地点
	RecordDate  *time.Time         `gorm:"column:recordDate" json:"recordDate,omitempty"`                              //记录日期
	Milestone   string             `gorm:"column:milestone" json:"milestone,omitempty"`                                //里程碑描述
	Mood        string             `gorm:"column:mood" json:"mood,omitempty"`                                          //心情
	Weather     string             `gorm:"column:weather" json:"weather,omitempty"`                                    //天气
	Tags        string             `gorm:"column:tags" json:"tags,omitempty"`                                          //标签
	ParentId    int32              `gorm:"column:parentId;NOT NULL" json:"parentId,omitempty"`                         //记录者ID
	Pinned      *bool              `gorm:"column:pinned;default:false;NOT NULL" json:"pinned,omitempty"`               //是否置顶
	ShowType    *int32             `gorm:"column:showType;default:1;NOT NULL" json:"showType,omitempty"`               //显示类型
	CreatedAt   *time.Time         `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time         `gorm:"column:updatedAt;NOT NULL" json:"updatedAt,omitempty"`

	Child       *Child             `gorm:"foreignKey:ChildId" json:"child,omitempty"`
	Parent      *User              `gorm:"foreignKey:ParentId" json:"parent,omitempty"`

	ImgConfigs  *[]map[string]any  `gorm:"-" json:"imgConfigs,omitempty"`
}

func (g *GrowthRecord) TableName() string {
	return "GrowthRecord"
}