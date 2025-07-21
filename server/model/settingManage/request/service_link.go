package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

var TableName = "service_link"

type ServiceLinkSearch struct {
	request.PageInfo
	Status *uint `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1停用）;"` //状态（0正常 1停用）
}

type ServiceLinkCreate struct {
	Link        string `json:"link" form:"link" binding:"required" gorm:"column:link;comment:链接;size:255;"`    //链接
	Image       string `json:"image" form:"image" binding:"required" gorm:"column:image;comment:图片;size:255;"` //图片
	Name        string `json:"name" form:"name" binding:"required" gorm:"column:name;comment:名字;size:64;"`     //名字
	Sort        int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                        //sort字段
	Type        uint   `json:"type" form:"type" gorm:"column:type;comment:类型;"`                                //类型
	Status      uint   `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1停用）;"`                 //状态（0正常 1停用）
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`            //createdTime字段
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`            //updatedTime字段
	CreatedUid  uint   `json:"createdUid" form:"createdUid" gorm:"column:created_uid;comment:创建者;size:19;"`    //创建者
}

func (ServiceLinkCreate) TableName() string {
	return TableName
}

type ServiceLinkUpdate struct {
	Id          int    `json:"id" form:"id" binding:"required" gorm:"primarykey;column:id;comment:id;size:20;"` //id
	Link        string `json:"link" form:"link" gorm:"column:link;comment:链接;size:255;"`                        //链接
	Image       string `json:"image" form:"image" gorm:"column:image;comment:图片;size:255;"`                     //图片
	Name        string `json:"name" form:"name" gorm:"column:name;comment:名字;size:64;"`                         //名字
	Sort        *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //sort字段
	Type        *uint  `json:"type" form:"type" gorm:"column:type;comment:类型;"`                                 //类型
	Status      *uint  `json:"status" form:"status" gorm:"column:status;comment:状态（0正常 1停用）;"`                  //状态（0正常 1停用）
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;;"`                     //updatedTime字段
	UpdatedUid  uint   `json:"updatedUid" form:"updatedUid" gorm:"column:updated_uid;comment:修改者;size:19;"`     //修改者
}

func (ServiceLinkUpdate) TableName() string {
	return TableName
}
