// 自动生成模板CouponIssued
package coupon

import (
	"gorm.io/gorm"
	"time"
)

// couponIssued表 结构体  CouponIssued
type CouponIssued struct {
	Id            *uint64            `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:20;"`                       //主键
	UserId        uint64             `json:"userId" form:"userId" gorm:"column:user_id;comment:归属用户ID;"`                         //归属用户ID
	CouponId      *uint64            `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券ID;"`                    //优惠券ID
	CouponName    string             `json:"couponName" form:"couponName" gorm:"column:coupon_name;comment:优惠券名称;size:50;"`      //优惠券名称
	PhoneNumber   string             `json:"phoneNumber" form:"phoneNumber" gorm:"column:phone_number;comment:手机号码;size:50;"`    //手机号码
	CouponAmount  *float64           `json:"couponAmount" form:"couponAmount" gorm:"column:coupon_amount;comment:优惠券金额;"`        //优惠券金额
	Type          *Type              `json:"type" form:"type" gorm:"column:type;comment:优惠券类型：0-未知券 1-注册送券 2-实名送券 3-人工送券;"`      //优惠券类型    ValidStart  *int `json:"validStart" form:"validStart" gorm:"column:valid_start;comment:有效期开始时间;size:19;"`  //有效期开始时间
	ValidStart    int64              `json:"validStart" form:"validStart" gorm:"column:valid_start;comment:有效期开始时间;size:10;"`    //有效期开始时间
	ValidEnd      int64              `json:"validEnd" form:"validEnd" gorm:"column:valid_end;comment:有效期开始时间;size:19;"`          //有效期开始时间
	UserType      uint               `json:"userType" form:"userType" gorm:"column:user_type;comment:有效期开始时间;size:19;"`          //有效期开始时间
	Status        IssuedCouponStatus `json:"status" form:"status" gorm:"column:status;comment:0未使用 1已使用 2已过期;"`                  //0未使用 1已使用 2已过期
	CreatedAt     int64              `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                      //createdAt字段
	CreatedAtTick int64              `json:"createdAtTick" form:"createdAtTick" gorm:"-"`                                        //createdAt字段
	UpdateAtTick  int64              `json:"updateAtTick" form:"updateAtTick" gorm:"-"`                                          //createdAt字段
	UpdatedAt     int64              `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime:milli column:updated_at;comment:;"` //updatedAt字段
}

// TableName couponIssued表 CouponIssued自定义表名 coupon_issued
func (c *CouponIssued) TableName() string {
	return "coupon_issued"
}

func (c *CouponIssued) AfterFind(*gorm.DB) error {
	if c.Status == NotUsed && time.Now().UnixMilli() > c.ValidEnd {
		c.Status = Expired
	}
	return nil
}
